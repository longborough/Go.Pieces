package main

import (
	"fmt"
	"math/big"
)

func results( inch chan int64 ) {     // Input for printing
	max := int64(0)                   // Initialise maximum value so far
	for next := <-inch; next > 0; next = <-inch { // Read until finished
		if (next > max) {             // A larger one?
			max = next                // Yes - update largest and 
			fmt.Printf("%v \n", next)   // Print largest so far
		}
	}
}

func makeprimes(
	next int64,
	pipe chan int64 ) {
	var i int64
	var p int64
	for i = 1; i < 10 ; i+=2 {
		p = next * 10 + i
		if ( big.NewInt(p).ProbablyPrime(10) ) {
			pipe <- p
			makeprimes(p,pipe)
		}
	}
}

func main() {
	var i int64
	ouch := make(chan int64)              // Create the output channel
	go results(ouch)                    // Start the printer
	for i = 1; i < 10 ; i+=1 {
		makeprimes(i,ouch)
		}
	close(ouch) // Send the done signal down the pipe
	fmt.Printf("Main done \n")
}
