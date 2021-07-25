package fundamentals

import "testing"

func Test_StructSample(t *testing.T) {
	want_name := "Bob"
	want_age := 100
	user := StructSample(want_name, want_age)

	if want_name != user.GetName() {
		t.Errorf("Expected %s, but actual %s", want_name, user.GetName())
	}

	if want_age != user.GetAge() {
		t.Errorf("Expected %d, but actual %d", want_age, user.GetAge())
	}
}
