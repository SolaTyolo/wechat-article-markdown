package main

import (
	"os"

	"github.com/kbsink-org/kbsink/cmd/kb-sink-md/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
