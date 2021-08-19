package pkg

import (
	"os"

	"github.com/jessevdk/go-flags"
)

var CMD = NoArgs

var Args struct {
	Paste  bool   `short:"p" long:"paste" description:"Adds context from clipboard"`
	File   string `short:"f" long:"file" description:"Adds the context from the given FILE" value-name:"FILE"`
	Dir    string `short:"d" long:"directory" description:"Adds the contexts from yaml files in the given DIR" value-name:"DIR"`
	Remove string `short:"r" long:"remove" description:"Removes CONTEXT-NAME from the config" value-name:"CONTEXT-NAME"`
	NS     string `short:"n" long:"namespace" description:"Sets the given NAMESPACE as the default ns" value-name:"NAMESPACE"`
}

const (
	NoArgs = iota
	Dir
	File
	Paste
	Remove
	NS
)

func init() {
	_, err := flags.ParseArgs(&Args, os.Args)
	if err != nil {
		// Don't exit during test runs
		if err.Error() != "unknown flag `t'" {
			os.Exit(1)
		}
	}

	defer ReadDefaultConfig()

	if len(os.Args) < 2 {
		CMD = NoArgs
		return
	}

	if Args.Dir != "" {
		CMD = Dir
		return
	}

	if Args.File != "" {
		CMD = File
		return
	}

	if Args.Paste {
		CMD = Paste
		return
	}

	if Args.Remove != "" {
		CMD = Remove
		return
	}

	if Args.NS != "" {
		CMD = NS
	}
}
