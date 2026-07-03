package utils

import (
	"encoding/json"
	"fmt"
)

func ToJson[T any](data T) string {
	json_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding json")
	}

	return string(json_data)
}

func JsonUnmarshal[T any](body []byte, data *T) (bool, error) {
	err := json.Unmarshal(body, data)
	if err != nil {
		return false, err
	}
	return true, nil
}