package main

import (
	"fmt"
	"time"
)

/*
Generators return the next value in a sequence each time they are called.
This means that each value is available as an output before the generator computes the next value.
Hence, this pattern is used to introduce parallelism in our program.
*/

func fibonacci(n int) chan int {
	mychannel := make(chan int)
	go func() {
		k := 0
		for i, j := 0, 1; k < n; k++ {
			mychannel <- i
			i, j = i+j, i
			time.Sleep(1 * time.Second)

		}
		close(mychannel)
	}()
	return mychannel
}

func main() {

	for i := range fibonacci(10) {
		//do anything with the nth term while the fibonacci()
		//is computing the next term
		fmt.Println(i)
	}
}
