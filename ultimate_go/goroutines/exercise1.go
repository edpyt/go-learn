package main

// Add imports.
import (
  "runtime"
  "sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		for i := 100; i > 0;  i-- {
		  println(i)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		for i := 0; i <= 100;  i++ {
		  println(i)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display "Terminating Program".
  println("\nTerminating Program")
}

