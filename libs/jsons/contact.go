package jsons

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact type
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON json data
var JSON = `{
  "name" : "Gopher",
  "title" : "Programmer",
  "contact" : {
    "home" : "081212312313",
    "cell" : "012124124121"
  }
}`

// ContactJSON contact json
func ContactJSON() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(c)
}
