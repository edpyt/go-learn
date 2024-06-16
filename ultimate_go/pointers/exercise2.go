// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.

// Declare a type named user.
type user struct {
	name string
	age  int
}
	
// Create a function that changes the value of one of the user fields.
func funcName(u *user, name string ){
	// Use the pointer to change the value that the
	// pointer points to.
  u.name = name
}

func main() {

	// Create a variable of type user and initialize each field.
  u := user{
	  name: "Bob",
	  age:  21,
	}

	// Display the value of the variable.
	println(u.name, u.age)

	// Share the variable, with the function you declared above.
	funcName(&u, "ed")

	// Display the value of the variable.
	println(u.name)
}

