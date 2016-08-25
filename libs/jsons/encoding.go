package jsons

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrettyEncoding display json in pretty
func PrettyEncoding() {
	c := make(map[string]interface{})
	c["name"] = "Eko Kurniawan"
	c["title"] = "No Title"
	c["contact"] = map[string]interface{}{
		"home": "09204124124",
		"cell": "23124214141",
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(data))
}

// UglyEncoding display json in pretty
func UglyEncoding() {
	c := make(map[string]interface{})
	c["name"] = "Eko Kurniawan"
	c["title"] = "No Title"
	c["contact"] = map[string]interface{}{
		"home": "09204124124",
		"cell": "23124214141",
	}

	data, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(data))
}

// Person type
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// StructEncoding encode struct data to json
func StructEncoding() {
	p := Person{
		FirstName: "Eko Kurniawan",
		LastName:  "Khannedy",
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(data))

}
