package maps

import "fmt"

// CreateEmptyMap create an empty map
func CreateEmptyMap() {
	slice := make([]int, 5)
	peta := make(map[int]int)

	fmt.Println(slice)
	fmt.Println(peta)
}

// InitializeMap initialize map
func InitializeMap() {
	peta := map[string]string{
		"first_name": "Eko Kurniawan",
		"last_name":  "Khannedy",
	}

	fmt.Println(peta)
}

// MapOfSlice create map that contains slice values
func MapOfSlice() {
	peta := map[string][]int{}

	peta["Eko"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	peta["Kurniawan"] = []int{2, 4, 6, 8, 0}
	peta["Khannedy"] = []int{1, 3, 5, 7, 9}

	fmt.Println(peta)
}

// RetrieveValue retrieve value from map
func RetrieveValue() {
	peta := map[string]string{
		"first_name":  "Eko",
		"middle_name": "Kurniawan",
		"last_name":   "Khannedy",
	}

	firstName, exists := peta["first_name"]
	if exists {
		fmt.Println(firstName)
	}
}

// IteratingMap iterating map with for
func IteratingMap() {
	peta := map[string]string{
		"first_name":  "Eko",
		"middle_name": "Kurniawan",
		"last_name":   "Khannedy",
	}

	for key, value := range peta {
		fmt.Println(key, value)
	}
}

// RemoveItemInMap remove item in map
func RemoveItemInMap() {
	peta := map[string]string{
		"first_name":  "Eko",
		"middle_name": "Kurniawan",
		"last_name":   "Khannedy",
	}

	delete(peta, "middle_name")

	fmt.Println(peta)
}

// removeMap removing data from map by key
func removeMap(peta map[string]string, key string) {
	delete(peta, key)
}

// MapString is map of string-string
type MapString map[string]string

// Remove remove key from map
func (peta MapString) Remove(key string) {
	removeMap(peta, key)
}

// Set set value in map
func (peta MapString) Set(key string, value string) {
	peta[key] = value
}

// PassingMap passing map is reference not value (copy data)
func PassingMap() {
	peta := MapString(map[string]string{})
	peta.Set("first_name", "Eko")
	peta.Set("middle_name", "Kurniawan")
	peta.Set("last_name", "Khannedy")

	fmt.Println(peta)

	peta.Remove("middle_name")
	fmt.Println(peta)
}
