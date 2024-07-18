// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic function named copyfy that is constrained to only
// making copies of slices of type string or int.
package main

import "fmt"

// Declare an interface named copyer that creates a constraint on
// string and int.
type copyer interface {
  string | int
}

// Implement a generic function named copyfy that accepts a slice of some
// type T but constrained on the copyer interface.
func copyfy[T copyer](c []T) []T {
  tmp :=  make([]T, len(c))
  copy(tmp, c)
  return tmp
}

func main() {

	// Construct a slice of string with three values.
	strList := []string{
	  "1", "2", "3",
	}

	// Call the copyfy function to make a copy of the slice.
	copyStrList := copyfy(strList)

	// Display the slice and the copy.
	fmt.Printf("%+v\n", copyStrList)

	// Construct a slice of int with three values.
	intList := []int{1, 2, 3,}

	// Call the copyfy function to make a copy of the slice.
	copyIntList := copyfy(intList)

	// Display the slice and the copy.
	fmt.Printf("%+v\n", intList)
	fmt.Printf("%+v\n", copyIntList)
}

