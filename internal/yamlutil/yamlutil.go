package yamlutil

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type YamlUtil struct {
}

func (y YamlUtil) ToObject(yamlStr string, any any) {

	err := yaml.Unmarshal([]byte(yamlStr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func (y YamlUtil) ToStructuredString(any any) string {
	jsonBytes, err := yaml.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func (y YamlUtil) FileContentToObject(filePaht string, any any) {
	yamlStr, err := os.ReadFile(filePaht)
	err = yaml.Unmarshal(yamlStr, &any)
	if err != nil {
		log.Fatal("convert yaml str failed", any)
	}
}

var YamlConverter = &YamlUtil{}
