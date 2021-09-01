package fundamentals

func SliceSample(count uint) []string {
	var strings_slice []string
	for i := uint(0); i < count; i++ {
		strings_slice = append(strings_slice, "John")
	}
	return strings_slice
}

func SliceSample2(length int) []int {
	numbers := []int{1, 2, 3, 4, 5, 6}
	return numbers[0:length]
}
