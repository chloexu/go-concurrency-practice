package main

import (
	// "fmt"
	"sync"
	// "time"
)

// unbuffered channel: receive (<-c) has to happen before the send (c <- 0). in other words, if receive is not ready, send is blocked
// var c = make(chan int)

// buffered channel: send (c <- 0) can happen before receive (<-c) is ready as long as buffer capacity is not reached
var c = make(chan int, 1)

var a string

func f() {
	a = "hello, world"
	<-c // block until receives from C
}

var l sync.Mutex

func foo() {
	a = "hello, world"
	l.Unlock()
}

var once sync.Once

func setup() {
	a = "hello, world"
}

func bar() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go bar()
	go bar()
}

func main() {
	/*
		// START OF TEST: sync lock & unlock
		l.Lock()
		go foo()
		l.Lock()
		print(a)
		// END OF TEST: sync lock & unlock
	*/

	// START OF TEST: sync once
	bar()
	// END OF TEST: sync once

	/*
		// START OF TEST: block with unbuffered vs. buffered channel
		// with unbuffered channel, c <- 0 is blocked until <-c in f() runs, so a is assigned with 'hello, world' first before the print, so print is hello, world.
		// with buffered channel, c <- 0 can happen almost at the time of f() and is not guaranteed to happen after value assignment of a, so print could be empty.
		go f()
		c <- 0
		print(a)
		// END OF TEST: block with unbuffered vs. buffered channel
	*/

	/*
		// START OF TEST: using buffered channel as rate limiter
		// below there are 6 different work functions defined, but because the channel limit is 3, only 3 jobs can run concurrently
		work := make([]func(a int) int, 6, 6) // I use interface{} to allow any kind of func
		work[0] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a + 1
			t := time.Now()
			fmt.Printf("%v----- plus 1 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		work[1] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a - 1
			t := time.Now()
			fmt.Printf("%v----- minus 1 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		work[2] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a + 2
			t := time.Now()
			fmt.Printf("%v----- plus 2 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		work[3] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a - 2
			t := time.Now()
			fmt.Printf("%v----- minus 2 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		work[4] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a + 3
			t := time.Now()
			fmt.Printf("%v----- plus 3 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		work[5] = func(a int) int {
			time.Sleep(3 * time.Second)
			r := a - 3
			t := time.Now()
			fmt.Printf("%v----- minus 3 is %v\n", t.Format("20060102150405"), r)
			return r
		} // good

		var limit = make(chan int, 3)
		a := 3
		for _, w := range work {
			go func(w func(a int) int) {
				limit <- 1
				w(a)
				<-limit
			}(w)
		}
		select {}
		// END OF TEST: using buffered channel as rate limiter
	*/
}
