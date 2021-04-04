package main

import (
	"fmt"
	"os"

	"./app"
)

const (
	// exit code if the program fails
	exitFail = 1
)

func main() {
	if err := app.Run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
