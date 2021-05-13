package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func main() {
	userPointer := &User{1, "user"}
	fmt.Println(userPointer.ID)
	userValue := User{1, "user"}
	fmt.Println(userValue.ID)
	OverwriteUserInfo(*userPointer, 3)
	fmt.Println(*userPointer)
	OverwriteUserInfo(userValue, 2)
	fmt.Println(userValue)
}

func OverwriteUserInfo(u User, newId int) {
	u.ID = newId
}
