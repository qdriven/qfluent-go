package json

import (
	"encoding/json"
	"log"
)

func ToStruct(jsonstr string, any any) {

	err := json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func ToJsonStr(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}
