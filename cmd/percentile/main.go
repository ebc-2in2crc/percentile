package main

import (
	"fmt"
	"os"

	"github.com/ebc-2in2crc/percentile"
)

func main() {
	cli := percentile.CLI{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	if err := cli.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
