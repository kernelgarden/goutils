package json

import "encoding/json"

type ArbitaryStruct map[string]interface{}

func FromJson(jsonString string) (ArbitaryStruct, error) {
	var f map[string]interface{}
	if err := json.Unmarshal(([]byte)(jsonString), &f); err != nil {
		return nil, err
	}

	return f, nil
}