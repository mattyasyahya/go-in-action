package jsons

import (
	"encoding/json"
	"fmt"
	"log"
)

// UnmarshalMap unmarshal json to map
func UnmarshalMap() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(c["name"])
	fmt.Println(c["title"])
	fmt.Println(c["contact"].(map[string]interface{})["home"])
	fmt.Println(c["contact"].(map[string]interface{})["cell"])
}
