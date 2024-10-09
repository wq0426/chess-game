package json

import (
	"github.com/goccy/go-json"
)

func FormatToBytes(data any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func ParseToAny(originData []byte) (any, error) {
	var data any
	err := json.Unmarshal(originData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
