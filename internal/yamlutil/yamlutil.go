package yamlutil

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ToObject(yamlStr string, any any) {

	err := yaml.Unmarshal([]byte(yamlStr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func ToYamlString(any any) string {
	jsonBytes, err := yaml.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func YamlFileToObject(yamlFile string, any any) {
	yamlStr, err := os.ReadFile(yamlFile)
	err = yaml.Unmarshal(yamlStr, &any)
	if err != nil {
		log.Fatal("convert yaml str failed", any)
	}
}
