package fundamentals

import "testing"

func Test_PointerSample(t *testing.T) {
	input := "new"
	want := "overwrite"
	PointerSample(&input)
	if want != input {
		t.Errorf("PointerSample = %s; should be %s", input, want)
	}
}
