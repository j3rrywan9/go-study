package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // Embedded type
	level string
}

func main() {
	adm := admin{
		user: user{
			name: "john smith",
			email: "john.smith@gmail.com",
		},
		level: "super",
	}
	
	// We can access the inner type's method directly
	adm.user.notify()

	// The inner type's method is promoted
	adm.notify()
}
