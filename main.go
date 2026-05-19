package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		if err := runScript(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "nachos-console: %v\n", err)
			os.Exit(1)
		}
		return
	}

	runShell()
}
