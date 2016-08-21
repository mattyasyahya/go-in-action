package slice

import "fmt"

// CreateMultidimentionSlice sample create multidimention slice
func CreateMultidimentionSlice() {
	slice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	fmt.Println(slice)
	slice[0] = append(slice[0], 2)
	fmt.Println(slice)
}
