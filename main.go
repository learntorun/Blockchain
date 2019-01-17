package main

import (
	"Blockchain-with-Golang/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
