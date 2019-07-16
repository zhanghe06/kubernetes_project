package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
)

func ShowJson(v interface{})  {
	// show json
	jsonIndentBytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println()
	}
	fmt.Println(string(jsonIndentBytes))
}

func ShowYaml(v interface{})  {
	// show yaml
	yamlBytes, err := yaml.Marshal(v)
	if err != nil {
		fmt.Println()
	}
	fmt.Println(string(yamlBytes))
}
