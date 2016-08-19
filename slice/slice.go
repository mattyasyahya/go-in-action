package slice

import "log"

// MakeSlice how to create a slice
func MakeSlice() {
	slice := make([]string, 5)
	log.Println(slice)

	slice2 := make([]string, 3, 5)
	log.Println(slice2)
}

// MakeWithLiteral make slice with literal
func MakeWithLiteral() {
	slice := []string{
		"eko", "kurniawan", "khannedy",
	}
	log.Println(slice)

	sliceInt := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}
	log.Println(sliceInt)

	sliceBig := []int{
		99: 100,
	}
	log.Println(sliceBig)
}

// EmptySlice create empty slice
func EmptySlice() {
	slice1 := make([]string, 0)
	log.Println(slice1)

	slice2 := []string{}
	log.Println(slice2)
}
