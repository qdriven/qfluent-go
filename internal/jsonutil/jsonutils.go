package jsonutil

import (
	"encoding/json"
	"log"
	"os"
)

type JsonUtil struct {
}

func (j JsonUtil) ToObject(jsonstr string, any any) {

	err := json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func (j JsonUtil) ToStructuredString(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func (j JsonUtil) FileContentToObject(filePath string, any any) {
	jsonstr, err := os.ReadFile(filePath)
	err = json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func NewJsonUtil() *JsonUtil {
	return &JsonUtil{}
}

var JsonConverter = NewJsonUtil()
