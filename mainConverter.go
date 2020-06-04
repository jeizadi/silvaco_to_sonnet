package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var fileLines []string
var file *os.File
var err error

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(dir)
	viper.Set("folder.directory", parent) // override to parent of current wd
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine) // take in command line args for folder path or first line
}

func main() {
	// read in files in current folder
	files, err := ioutil.ReadDir(viper.GetString("folder.directory"))
	if err != nil {
		log.Fatal(err)
	}

	// make a new directory with the current date, if duplicate then rename indicating a different folder
	date := time.Now()
	tempSonnetFolder := viper.GetString("folder.directory") + "/" + date.Format("2006-01-02")
	finalSonnetFolder := tempSonnetFolder
	exist, err := exists(finalSonnetFolder)
	if err != nil {
		log.Fatal(err)
	}
	for counter := 1; exist == true; counter++ { // if there already exists a folder by the same name, rename
		finalSonnetFolder = tempSonnetFolder + "(" + strconv.Itoa(counter) + ")"
		exist, _ = exists(finalSonnetFolder)
	}

	os.Mkdir(finalSonnetFolder, os.ModePerm) // create the new folder

	// convert sonnet file to silvaco
	for _, file := range files {
		if strings.Contains(file.Name(), ".log") { // check if the file is a .log type or Silvaco file
			filePath := viper.GetString("folder.directory") + "/" + file.Name()
			currentFile, err := os.Open(filePath) // open file at current file path
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}

			// read in lines from file and store
			scanner := bufio.NewScanner(currentFile)
			fileLines = nil // reset global variable fileLines
			for scanner.Scan() {
				fileText := scanner.Text()
				if fileText[0:1] == "d" { // isolate only lines starting with d and add them to the files lines var
					fileLines = append(fileLines, fileText)
				}
			}

			// after we are done reading in the lines from the file, close and move it into the new folder
			newPathFile := finalSonnetFolder + "/" + file.Name()
			currentFile.Close()
			os.Rename(filePath, newPathFile)

			// write to newly created Sonnet file
			// remove the .log and add an .s2p extension for new Sonnet file
			newFileName := strings.Trim(file.Name(), ".log")
			f, err := os.Create(finalSonnetFolder + "/" + newFileName + ".s2p")
			if err != nil {
				fmt.Println(err)
				return
			}
			l, err := f.WriteString(format(fileLines)) // format the file and write necessary file lines to it
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "bytes written successfully") // new file successfully created and written to
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// format takes in the Silvaco file lines as a slice of string and converts them to Sonnet through elementary operations
func format(unformatted []string) string {
	uString := unformatted
	finalString := viper.GetString("first.line") // added as the first line to every file

	// formatting changes for Silvaco to Sonnet
	for _, v := range uString {
		finalString += "\n"
		v = strings.Trim(v, "d ")
		tempString := strings.Split(v, " ")
		tempString = tempString[0:9]
		for _, v := range tempString {
			finalString = finalString + v + " "
		}
	}
	return finalString
}

// exists checks whether a folder was already created on the same date and if so, it increments the naming convention by one each time
// 2006-01-02 becomes 2006-01-02(1) becomes 2006-01-02(2)...
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
