package fundamentals

import "time"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	BirthDate   time.Time
}

func StructSample(name string, age int) User {
	user := User{
		Name:        name,
		Age:         age,
		PhoneNumber: "xxx-yyy-zzzz",
		BirthDate:   time.Date(1900, time.April, 1, 0, 0, 0, 0, time.Local),
	}
	return user
}
