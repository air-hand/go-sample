package fundamentals

import (
	"testing"
)

func Test_SliceSample(t *testing.T) {
	input := 100
	slice := SliceSample(uint(input))
	if input != len(slice) {
		t.Logf("Expected %d, but Actual %d", input, len(slice))
		t.FailNow()
	}
}

func Test_SliceSample2(t *testing.T) {
	numbers := SliceSample2(1)
	if 1 != len(numbers) {
		t.Errorf("Expected 1, Actual %d", len(numbers))
	}
}
