package main

import (
	"os"

	"github.com/tensor-programming/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	
	//os.Setenv("NODE_ID", "7000") //Can uncomment for testing if you keep receiving <NODE_ID env is not set> errors

	cmd := cli.CommandLine{}
	cmd.Run()
}
