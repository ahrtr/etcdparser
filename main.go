package main

import (
	"os"
	"fmt"
	"github.com/ahrtr/etcdparser/cmd"

)

func main()  {
	root := cmd.CreateRootCommand()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println()
}

