package main

import (
	"github.com/hIMEI29A/kilo"
	"os"
)

func main() {
	kilo.EnableRawMode()
	defer kilo.DisableRawMode()
	kilo.InitEditor()
	if len(os.Args) > 1 {
		kilo.EditorOpen(os.Args[1])
	}

	kilo.EditorSetStatusMessage("HELP: Ctrl-S = save | Ctrl-Q = quit | Ctrl-F = find")

	for {
		kilo.EditorRefreshScreen()
		kilo.EditorProcessKeypress()
	}
}
