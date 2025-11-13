package main

func main() {
	//all the cases are of deadlock
	/*
		ch := make(chan int)
		ch <- 10
		fmt.Println(<-ch)
	*/
	/*
		ch := make(chan int)
		fmt.Println(<-ch)
		ch <- 10
	*/

	/*

		// receiver always blocks until the sender sends the value
		   fmt.Println(<-ch)

			wg.Go(func() {

				ch <- 10
			})


	*/
}
