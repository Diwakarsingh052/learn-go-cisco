package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

// context allow us to cancel the work if it does not finish in provided time
func main() {
	// Background context is empty context which is used to store timeout values
	// It is never canceled, has no values, and has no deadline.
	//It is typically used by the main function, initialization,
	//and tests, and as the top-level Context for incoming requests.
	ctx := context.Background()

	//updating the context with timeout
	// adding a timer value to the context
	// WithTimeout is a func which updates the existing context with timeout value
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// cancel clean up the resources taken up by the context
	// if cancel is not called in defer , ctx would be immediately canceled
	// which means the timer would be closed and the resources would be cleaned up

	//ctx := context.Background()
	//sql.DB.ExecContext(ctx, "SELECT * FROM users")
	//sql.DB.QueryContext()

	// constructing a new request with context
	// we have not made the request yet to the remote server
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// doing the request to the remote server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	// reading the response body in bytes
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// printing the response body
	log.Println(string(bytes))

}
