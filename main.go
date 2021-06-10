// Copyright (c) 2021, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

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
