package fundamentals

import "fmt"

func StringSample(input string) (string, error) {
	var str string
	str = "Goodbye, cruel world (%s)."

	return fmt.Sprintf(str, input), nil
}
