// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic map type.
package main

// Declare a generic type named keymap that uses an underlying type of map
// with a key of type string and a value of some type T.
type keymap[T any] map[string]T

// Implement a method named set that accepts a key of type string and a value
// of type T.
func (kv keymap[T]) set(k string, v T) {
  kv[k] = v
}

// Implement a method named get that accepts a key of type string and return
// a value of type T and true or false if the key is found.
func (kv keymap[T]) get(k string) (T, bool) {
  v, ok := kv[k]
  return v, ok
}

func main() {
	// Construct a value of type keymap that stores integers.
	var intKeyMap keymap[int] = map[string]int{}

	// Add a value with key "jack" and a value of 10.
	intKeyMap.set("jack", 10)

	// Add a value with key "jill" and a value of 20.
	intKeyMap.set("jill", 20)

	// Get the value for "jack" and verify it's found.
	jack, ok := intKeyMap.get("jack")

	if !ok {
	  println("Error, jack not found")
	  return
	}

	// Print the value for the key "jack".
	println(jack)

	// Get the value for "jill" and verify it's found.
	jill, ok := intKeyMap.get("jill")

	if !ok {
	  println("error, jill not found")
	  return
	}

	// Print the value for the key "jill".
	println(jill)
}
