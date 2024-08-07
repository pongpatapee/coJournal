package utils

import (
	"encoding/json"
	"fmt"
)

func PrintObject(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("error trying to print object")
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
