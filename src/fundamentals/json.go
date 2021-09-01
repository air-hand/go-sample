package fundamentals

import (
	"encoding/json"
	"fmt"
	"log"
)

type QueryString struct {
	Query        string `json:"query"`
	DefaultField string `json:"default_field"`
}

type Query struct {
	QueryString QueryString `json:"query_string"`
}

type QueryDSL struct {
	Query Query `json:"query"`
}

/*
// You can declare a nested struct like below, but initialize with hard-code it *too* hard.
type QueryDSL struct {
	Query struct {
		QueryString struct {
			Query        string `json:"query"`
			DefaultField string `json:"default_field"`
		} `json:"query_string"`
	} `json:"query"`
}
*/

func JsonToQuery(default_field string) QueryDSL {
	query_json := fmt.Sprintf(`
{
    "query": {
        "query_string": {
            "query": "(foo) OR (bar)",
            "default_field": "%s"
        }
    }
}
`, default_field)
	var unmarshalled QueryDSL
	err := json.Unmarshal([]byte(query_json), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling", err)
	}
	return unmarshalled
}

func QueryToJson(query_string string) string {
	query := QueryDSL{
		Query: Query{
			QueryString: QueryString{
				Query:        query_string,
				DefaultField: "content",
			},
		},
	}
	query_json, err := json.MarshalIndent(query, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	return string(query_json)
}
