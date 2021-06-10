// Copyright (c) 2021, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package cmd

import (
	"encoding/json"
	"fmt"
)

func formatStructInJSON(val interface{}, rawFormat bool) (string, error) {
	var (
		formattedData []byte
		err           error
	)

	if rawFormat {
		formattedData, err = json.Marshal(val)
	} else {
		formattedData, err = json.MarshalIndent(val, "", "    ")
	}
	if err != nil {
		return "", err
	}

	return string(formattedData), nil
}

func printJsonObject(header, data string) {
	if len(header) > 0 {
		fmt.Printf("%s: \n%s\n", header, data)
	} else {
		fmt.Printf("%s\n", data)
	}
}

func printSeparator() {
	fmt.Println()
	fmt.Println("-----------------------------------------------------")
	fmt.Println()
}
