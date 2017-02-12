package main

import (
	"fmt"
	"math/big"
)

func results( inch chan *Int ) {      // Input for printing
	var max *Int                      // Initialise maximum value so far
	for next := <-inch; next > 0; next = <-inch { // Read until finished
		if (next > max) {             // A larger one?
			max = next                // Yes - update largest and 
			fmt.Printf("%v \n", next)   // Print largest so far
		}
	}
}

func makeprimes(
	next *Int,
	mult *Int,
	pipe chan *Int ) {
	var i *Int
	var p *Int
//	if ( mult >= 1000000000000000000 ) {
//		return
//	}
//	fmt.Printf("%v %v\n", next, mult)
//	fmt.Printf("%v\n", next)
	for i = 0; i < 10 ; i++ {
		p = i * mult + next
		if ( big.NewInt(p).ProbablyPrime(10) ) {
			pipe <- p
			makeprimes(p,mult*10,pipe)
		}
	}
}

func main() {
	var i int64
	ouch := make(chan int64)              // Create the output channel
	go results(ouch)                    // Start the printer
	for i = 3; i < 10 ; i++ {
		makeprimes(i,1,ouch)
		}
	close(ouch) // Send the done signal down the pipe
	fmt.Printf("Main done \n")
}
