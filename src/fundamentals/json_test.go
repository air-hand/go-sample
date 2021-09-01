package fundamentals

import (
	"strings"
	"testing"
)

func Test_QueryToJson(t *testing.T) {
	expected_included := "search word"
	json_value := QueryToJson(expected_included)
	if !strings.Contains(json_value, expected_included) {
		t.Errorf("%s must include %s", json_value, expected_included)
	}
}

func Test_JsonToQuery(t *testing.T) {
	expected := "field_to_search"
	query := JsonToQuery(expected)
	if query.Query.QueryString.DefaultField != expected {
		t.Errorf("Expected %s, but actual %s", expected, query.Query.QueryString.DefaultField)
	}
}
