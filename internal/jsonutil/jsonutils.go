package jsonutil

import (
	"encoding/json"
	"log"
	"os"
)

func ToObject(jsonstr string, any any) {

	err := json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func ToJsonString(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func ToJsonStringFromFile(jsonFile string, any any) {
	jsonstr, err := os.ReadFile(jsonFile)
	err = json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}
