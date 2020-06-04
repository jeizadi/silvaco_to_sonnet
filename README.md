# Silvaco(.log) to Sonnet(.s2p) File Conversion 
Program to convert Silvaco (.log) type files to Sonnet (.s2p) type files.

# How to use this program
To use this program you will need to build an executable for your software preference and place this binary file into the folder with your .log files that you wish to use for your Silvaco (.log) to Sonnet (.log) conversion. 
To run the algorithm, all you will need to do is double tap to run the binary file and your exisiting .log files will be converted into .s2p files and moved into a new folder named according to date. The files will have the same naming convention but you will see both a .log and .s2p file for each file name. 

# Modifications
You are able to change the first line of the new Sonnet(.s2p) file from default "# Hz S RI R 50" usin the tag --first.line.
```bash
--first.line string
```
You are able to manually input the folder directory of the Silvaco (.log) file to convert with the tag --folder.directory.
```bash
--folder.directory string
```

# Build
Build an executable for Mac, Windows, and Linux OS.
```bash
make build_all
```
Build an executable for Mac OS.
```bash
make build_mac
```
Build an executable for Windows OS.
```bash
make build_win
```
Build an executable for Linux OS.
```bash
make build_linux
```
