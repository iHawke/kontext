package main

import (
	"github.com/ihawke/kontext/pkg"
)

func main() {

	switch pkg.CMD {
	case pkg.NoArgs:
		pkg.NoArgsRun()
	case pkg.Dir:
		pkg.Directory()
	case pkg.File:
		pkg.OneFile()
	case pkg.Paste:
		pkg.PasteFromClipboard()
	case pkg.Remove:
		pkg.RemoveContext()
	case pkg.NS:
		pkg.AddNamespace()
	}
}
