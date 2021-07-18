package fundamentals

import (
	"regexp"
	"testing"
)

func Test_StringSample(t *testing.T) {
	input := "something"
	want := regexp.MustCompile(`\b` + input + `\b`)
	got, err := StringSample(input)
	if !want.MatchString(got) || err != nil {
		t.Fatalf("StringSample = %s; want to match with %s", got, input)
	}
}
