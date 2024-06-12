// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {
	var	x, y, z int
	var	p, o, s string
	var b bool

	fmt.Println(x, y, z)
	fmt.Println(p, o, s)
	fmt.Println(b)

	var f float32
	f = float32(3.14)
	fmt.Println(f)
}

