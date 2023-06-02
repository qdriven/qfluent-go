package jsonutil

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Address struct {
	Location string `json:"location"`
	ZipCode  string `json:"zipCode"`
}

/*
*  What's different between & and *
 */
func TestToJsonString(t *testing.T) {
	addr := Address{Location: "test", ZipCode: "109032"}
	fmt.Println(addr.Location)
	person := &Person{
		Name:    "Person",
		Age:     10,
		Address: addr,
	}
	result := ToJsonString(&addr)
	fmt.Println(result)
	personStr := ToJsonString(&person)
	fmt.Println(personStr)
}

func TestToStruct(t *testing.T) {
	jsonStr := "{\"name\":\"Person\",\"age\":10,\"address\":{\"location\":\"test\",\"zipCode\":\"109032\"}}\n"
	p := &Person{}
	ToObject(jsonStr, p)
	fmt.Println(p.Age, p.Name, p.Address.ZipCode, p.Address.Location)

}
