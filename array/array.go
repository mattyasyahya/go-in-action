/*
Package array contains sample array utilities
*/
package array

import "log"

// DeclaringArray sample declaring array
func DeclaringArray() {
	var array [5]int
	array[0] = 1
	log.Println(array)
}

// DeclaringArrayUsingArrayLiteral sample declaring array using array literal
func DeclaringArrayUsingArrayLiteral() {
	array := [5]int{
		1, 2, 3, 4, 5,
	}
	log.Println(array)
}

// CalculatingSize declaring array with calculating size
func CalculatingSize() {
	array := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}
	log.Println(array)
}

// SpecificElement declaring array with specific element
func SpecificElement() {
	array := [5]int{
		1: 10,
		3: 12,
	}
	log.Println(array)
}

// AccessingArray accesing array element
func AccessingArray() {
	array := [5]int{1, 2, 3, 4, 5}
	log.Println(array[2])
	log.Println(array[3])
}

// Pointer create array pointer
func Pointer() {
	array := [5]*int{1: new(int), 3: new(int)}
	*array[1] = 10
	*array[3] = 10
	log.Println(array)
}

// AssigningArray sample assigning array
func AssigningArray() {
	var array1 [5]int
	array2 := [5]int{1, 2, 3, 4, 5}

	array1 = array2

	array1[0] = 5
	array2[0] = 10

	log.Println(array1)
	log.Println(array2)
}

// AssigningArrayPointer sample assigning array pointer
func AssigningArrayPointer() {
	var array1 [5]*string
	array2 := [5]*string{new(string), new(string), new(string), new(string), new(string)}

	array1 = array2

	*array2[0] = "eko"
	*array2[1] = "kurniawan"
	*array2[2] = "khannedy"
	*array2[3] = "keren"
	*array2[4] = "banget"

	log.Println(array1)
	log.Println(array2)
}

// TwoDimentionalArray sample array two dimentional
func TwoDimentionalArray() {
	array := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	array[0][0] = 0
	array[0][1] = 1

	log.Println(array)
}
