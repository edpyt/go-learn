// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
  name string
  email string
  age int
}

func main() {
	// Declare variable of type user and init using a struct literal.
	u := user{
	  name: "bob",
	  email: "edrsa1321s@gmail.com",
	  age: 21,
	}

	// Display the field values.
  fmt.Println("Name", u.name)
  fmt.Println("email", u.email)
  fmt.Println("age", u.age)

	// Declare a variable using an anonymous struct.
	u1 := struct {
    name string
    email string
    age int
  } {
    name: "bib",
    email: "bibobbe13l@mail.com",
    age: 19,
  }

	// Display the field values.
  fmt.Println("Name", u1.name)
  fmt.Println("email", u1.email)
  fmt.Println("age", u1.age)
}

