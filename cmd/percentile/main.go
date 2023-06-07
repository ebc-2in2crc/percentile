package main

import (
	"fmt"
	"github.com/ebc-2in2crc/percentile"
	"os"
)

func main() {
	cli := percentile.CLI{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	if err := cli.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
