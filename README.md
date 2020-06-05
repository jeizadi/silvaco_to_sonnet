# Silvaco(.log) to Sonnet(.s2p) File Conversion 
Silvaco_to_Sonnet is a program made to convert Silvaco (.log) type files to Sonnet (.s2p) type files.

# Contribution
Created by Jenna Eizadi for the Hai Lab at the University of Wisconsin-Madison. For more information or inquiries visit the <a href="https://hailab.wisc.edu/Contact.html">Hai Lab</a> website or contact us at eizadi@wisc.edu

# Download
Build an executable for your software preference (Mac, Windows or Linux).

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
# Installation
Place the executable into a folder and place this new folder into the folder containing your files with the .log extension that you wish to convert from Silvaco (.log) to Sonnet (.log). <br>
<i> Please note that you must include the .log extension on all files you wish to convert from .log to .s2p </i> 

# Reformat your data
Run the executable to convert all .log into .s2p files. Both the existing .log files and new corresponding .s2p files will be moved into a new folder named according to today's date. <br>
<i> If multiple conversions on the same date, the naming convention will vary (ex. 01/01/01 and 01/01/01 (1)) </i>

# Modifications
To change the first line of the new Sonnet(.s2p) file from default "# Hz S RI R 50", use the tag --first.line.
```bash
--first.line string
```
To manually input the folder directory of the Silvaco (.log) files to convert, use the tag --folder.directory.
```bash
--folder.directory string
```
