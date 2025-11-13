package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	v, err := SlowFunc(ctx, "10")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(v)

}

// error should be the last return value returned from the function
// whenever error happens set other values to their default values

//ctx should be the first argument of the function

func SlowFunc(ctx context.Context, s string) (int, error) {

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	time.Sleep(1 * time.Second)
	select {
	// Done channel would receive a signal when the context is canceled or times is over
	case <-ctx.Done():
		fmt.Println("context is done")
		return 0, ctx.Err()

		// if none cases are ready, default case would be executed
	default:
		fmt.Println("time is not up")
	}
	return i, nil
}
