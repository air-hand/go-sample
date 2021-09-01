package fundamentals

import "log"

func ifElse() {
	if true == false && 0 != -1 {
		log.Println("foo")
	} else if "dog" == "cat" {
		log.Println("bar")
	} else {
		log.Println("baz")
	}
}

func switchCase() {
	var_ := "foo"

	switch var_ {
	case "foo":
		log.Printf("%s\n", var_)
	case "bar":
		log.Printf("%s is bar\n", var_)
	default:
		log.Printf("default: %s", var_)
	}
}

func forLoop() {
	for i := 0; i <= 10; i++ {
		log.Printf("i: %d. count to 10.", i)
	}

	strings := []string{"foo", "bar", "baz"}
	for _, str := range strings {
		log.Printf("str: %s", str)
	}

	strings_map := make(map[string]string)
	strings_map["hoge"] = "foo"
	strings_map["fuga"] = "bar"
	strings_map["piyo"] = "baz"
	for key, value := range strings_map {
		log.Printf("%s %s", key, value)
	}

	var sentence = "Goodbye cruel world."
	sentence = "f"
	for i, c := range sentence {
		log.Printf("%d %c", i, c)
	}
}
