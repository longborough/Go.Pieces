// esieve implements a Sieve of Eratosthenes
// as a series of channels connected together
// by goroutines
package main

import "fmt"

func sieve(mine int, // This instance's own prime
	inch chan int, // Input channel from lower primes
	done chan int, // Channel for signalling shutdown
	count int) { // Number of primes - counter
	start := true                                 // First-number switch
	ouch := make(chan int)                        // Output channel, this instance
	fmt.Printf("%v ", mine)                       // Print this instance's prime
	for next := <-inch; next > 0; next = <-inch { // Read input channel
		if (next % mine) > 0 { // Divisible by my prime?
			if start { // No; first time through?
				go sieve(next, ouch, done, count+1) // First number,
				// create instance for it
				start = false // First time done
			} else { // Not first time
				ouch <- next // Pass to next instance
			}
		}
	}
	if start { // Just starting?
		close(done)                     // Yes - we're last in pipe - signal done
		print("\n", count, " primes\n") // Number of primes/goroutines
	} else {
		close(ouch) // No - send the signal down the pipe
	}
}

func main() {
	lim := 100000                             // Let's do up to lim
	done := make(chan int)                    // Create the done return channel
	ouch := make(chan int)                    // Create the first segment of the pipe
	go sieve(2, ouch, done, 1)                // Create the first instance for '2'
	for prime := 3; prime < lim; prime += 1 { // Generate odd numbers
		ouch <- prime // Send numbers down the pipe
	}
	close(ouch) // Send the done signal down the pipe
	<-done      // and wait for it to come back
}
