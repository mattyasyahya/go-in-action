package slice

import (
	"fmt"
	"log"
)

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

// ChangeSlice change slice data
func ChangeSlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	log.Println(slice1)
	slice1[0] = slice1[0] * 10
	log.Println(slice1)
}

// TakeSliceAsSlice take slice as slice
func TakeSliceAsSlice() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	log.Println(slice1)
	slice2 := slice1[5:len(slice1)]
	log.Println(slice2)
}

// ChangeArraySlice change the array of 2 slices
func ChangeArraySlice() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	log.Println(slice1)
	slice2 := slice1[5:len(slice1)]
	log.Println(slice2)

	slice1[6] = slice1[6] * 10

	log.Println(slice1)
	log.Println(slice2)
}

// AppendSlice append new data to slice
func AppendSlice() {
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := slice1[3:8]

	log.Println(slice1)
	log.Println(slice2)

	slice2 = append(slice2, 100)

	log.Println(slice1)
	log.Println(slice2)
}

// AppendResizeSlice append and resize the slice
func AppendResizeSlice() {
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(slice1)
	slice2 := slice1[5:len(slice1)]
	log.Println(slice2)

	slice2 = append(slice2, 10)
	log.Println(slice1)
	log.Println(slice2)

	slice2[0] = 100
	log.Println(slice1)
	log.Println(slice2)
}

// CapacitySet set slice capacity
func CapacitySet() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	slice := source[2:3:3]

	// this will not edit source array because the capacity is until index 3
	slice = append(slice, "Kiwi")

	log.Println(source)
	log.Println(slice)
}

// AppendTwoSlice append two slice
func AppendTwoSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	result := append(slice1, slice2...)

	log.Println(result)
}

// IteratingSlice iterating slice or array
func IteratingSlice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, value := range slice {
		log.Println(value)
	}
}

// IteratingSliceIsCopy iterating slice is copy the value,
// not reference to the value
func IteratingSliceIsCopy() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, value := range slice {
		fmt.Printf("index %d with ref %d, real ref %d \n", i, &value, &slice[i])
	}
}
