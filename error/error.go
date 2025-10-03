package main

import (
	"fmt"
	"strings"
)

type error interface {
	Error() string
}

type user struct {
	Name, email string
}

type errorValidation struct {
	Field, Message string
}

// use interface
func (e errorValidation) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Message)
}

// custom error
func userValidation(name, email string) error {
	if name == "" {
		return errorValidation{
			Field:   name,
			Message: "Name is Empty",
		}
	}

	if len(name) < 6 {
		return errorValidation{
			Field:   name,
			Message: "User Name Length less than 6",
		}
	}

	if email == "" {
		return errorValidation{
			Field:   email,
			Message: "Email is Empty",
		}
	}

	if !strings.Contains(email, "@") {
		return errorValidation{
			Field:   email,
			Message: "Email is must contain @",
		}
	}

	return nil
}
func main() {
	users := []user{}
	user1 := user{"", "weslin@gmail.com"}
	user2 := user{"shabana", "shabana@gmail.com"}
	user3 := user{"shalini", "shalinigmail.com"}
	user4 := user{"ramkumar", "ramkumar@gmail.com"}

	users = append(users, user1, user2, user3, user4)

	for _, userdetails := range users {
		fmt.Println(userdetails)
		err := userValidation(userdetails.Name, userdetails.email)
		if err != nil {
			if value, ok := err.(errorValidation); ok {
				fmt.Println(value.Field, value.Message)
			} else {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Valid User")
		}
	}
}
