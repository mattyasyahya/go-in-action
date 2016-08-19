package math

// Average get average from slice
func Average(data ...int) float64 {
	var total float64
	for _, value := range data {
		total += float64(value)
	}
	return total / float64(len(data))
}
