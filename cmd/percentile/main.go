package main

import (
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	cli := CLI{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	if err := cli.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
