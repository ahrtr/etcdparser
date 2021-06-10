package main

import (
	"fmt"
	"github.com/ahrtr/etcdparser/cmd"
	"os"
)

func main() {
	root := cmd.CreateRootCommand()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println()
}
