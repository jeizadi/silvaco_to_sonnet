package main

import (
	"github.com/spf13/pflag"
)

const (
	// configuration defaults
	defaultFolderDirectory = "/" //overridden in init routine with parent of current wd
	defaultFirstLine       = "# Hz S RI R 50"
)

var (
	// define flag overrides
	flagFilePath  = pflag.String("folder.directory", defaultFolderDirectory, "path of Silvaco file")
	flagFirstLine = pflag.String("first.line", defaultFirstLine, "first line of Sonnet file")
)
