// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Use the template and follow the directions. You will be writing a web handler
// that performs a mock database call but will timeout based on a context if the call
// takes too long. You will also save state into the context.
package main

// Add imports.
import (
  "context"
  "encoding/json"
  "log"
  "net/http"
  "time"
)

// Declare a new type named `key` that is based on an int.
type key int

// Declare a constant named `userIPKey` of type `key` set to
// the value of 0.
const userIPKey key = 0

// Declare a struct type named `User` with two `string` based
// fields named `Name` and `Email`.
type User struct {
  Name string
  Email string
}

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// routes sets the routes for the web service.
func routes() {
	http.HandleFunc("/user", findUser)
}

// Implement the findUser function to leverage the context for
// both timeouts and state.
func findUser(rw http.ResponseWriter, r *http.Request) {

	// Create a context that timeouts in fifty milliseconds.
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)

	// Defer the call to cancel.
	defer cancel()

	// Save the `r.RemoteAddr` value in the context using `userIPKey`
	// as the key. This call returns a new context so replace the
	// current `ctx` value with this new one. The original context is
	// the parent context for this new child context.
	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	// Create a channel with a buffer size of 1 that works with
	// pointers of type `User` 
	ch := make(chan *User, 1)

	// Use this goroutine to make the database call. Use the channel
	// to get the user back.
	go func() {

		// Get the `r.RemoteAddr` value from the context and log
		// the value you get back.
		if ip, ok := ctx.Value(userIPKey).(string); ok {
		  log.Println("Start DB for IP", ip)
		}

		// Call the `readDatabase` function provided below and
		// send the returned `User` pointer on the channel.
		ch <- readDatabase()

		// Log that the goroutine is terminating.
		log.Println("DB goroutine terminated")
	}()

	// Wait for the database call to finish or the timeout.
	select {

	// Add a case to wait on the channel for the `User` pointer.
	case u := <-ch:
	  sendResponse(rw, u, http.StatusOK)
	  return
	case <-ctx.Done():
	  e := struct{ Error string }{ctx.Err().Error()}
	  sendResponse(rw,e,http.StatusRequestTimeout)
	  log.Println("Sent StatusRequestTimeout")
	  return
	}
}

// readDatabase performs a pretend database call with
// a second of latency.
func readDatabase() *User {
	u := User{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	// Create 100 milliseconds of latency.
	time.Sleep(100 * time.Millisecond)

	return &u
}

// sendResponse marshals the provided value into json and returns
// that back to the caller.
func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}

