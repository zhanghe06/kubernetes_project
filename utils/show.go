package utils

import (
	"encoding/json"
	"fmt"
)

func ShowJson(v interface{})  {
	// show
	jsonIndentBytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println()
	}
	fmt.Println(string(jsonIndentBytes))
}
