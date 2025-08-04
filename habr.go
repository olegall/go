package main

// fib returns a channel which transports fibonacci numbers
func fib(length int) <-chan int {
	// make buffered channel
	c := make(chan int, length)

	// run generation concurrently
	go func() { // убрать горутину - обёртку - результат тот же
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	// return channel
	return c
}
