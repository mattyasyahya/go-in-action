package slice

import "fmt"

// CreateMultidimentionSlice sample create multidimention slice
func CreateMultidimentionSlice() {
	slice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	fmt.Println(slice)
	slice[0] = append(slice[0], 2)
	fmt.Println(slice)
}

func changeSliceToZero(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
}

// PassingSliceToFunction passing slice to function
func PassingSliceToFunction() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(slice)
	changeSliceToZero(slice)
	fmt.Println(slice)
}
