package jsons

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Location type
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Viewport type
type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// Geometry type
type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

// Photo type
type Photo struct {
	Height         int      `json:"height"`
	Width          int      `json:"width"`
	PhotoReference string   `json:"photo_reference"`
	HTMLAttributes []string `json:"html_attributions"`
}

// Result json
type Result struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
	Icon             string   `json:"icon"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Photos           []Photo  `json:"photos"`
	PlaceID          string   `json:"place_id"`
	Reference        string   `json:"reference"`
	Types            []string `json:"types"`
}

// Response json
type Response struct {
	Results        []Result `json:"results"`
	Status         string   `json:"status"`
	HTMLAttributes []string `json:"html_attributions"`
}

// GoogleMapsPlaceSearch sample parsing json
func GoogleMapsPlaceSearch() {
	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?key=AIzaSyDJJB_32448J6EKl8EeQpKQS37AlY2jlDk&query=Bandung,%20Indonesia"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	// _, err = io.Copy(os.Stdout, resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(response)
}
