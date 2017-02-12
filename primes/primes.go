package main

import (
	"fmt"
	"math/big"
)

func results( inch chan int64 ) {     // Input for printing
	max := int64(0)                   // Initialise maximum value so far
	count := int64(0)
	for next := <-inch; next > 0; next = <-inch { // Read until finished
		count ++
		if (next > max) {             // A larger one?
			max = next * 0            // Yes - update largest and 
			fmt.Printf("%v %v\n", next, count) // Print largest so far
		}
	}
}

func makeprimes(
	next int64,
	mult int64,
	pipe chan int64 ) {
	var i int64
	var p int64
	if ( mult >= 9000000000000000000 ) {
		return
	}
	for i = 1; i < 10 ; i++ {
		p = i * mult + next
		if ( big.NewInt(p).ProbablyPrime(10) ) {
			pipe <- p
			makeprimes(p,mult*10,pipe)
		}
	}
//	if ( mult < 1 * next ) {
//		makeprimes(next,mult*10,pipe)
//	}
}

func main() {
	var i int64
	ouch := make(chan int64)              // Create the output channel
	go results(ouch)                    // Start the printer
	for i = 3; i < 10 ; i+=2 {
		makeprimes(i,1,ouch)
		}
	close(ouch) // Send the done signal down the pipe
	fmt.Printf("Main done \n")
}
