package fundamentals

import "testing"

func Test_MapSample(t *testing.T) {
	m := MapSample()
	john_doe := m["John Doe"]
	expected := 100000
	if john_doe.GetAge() != expected {
		t.Logf("Expected %d, but Actual %d", expected, john_doe.GetAge())
		t.FailNow()
		return
	}
}
