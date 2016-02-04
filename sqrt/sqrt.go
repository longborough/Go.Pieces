// Sqrt calculates integer square roots and residues by shift-subtract
package main

/*
import 	(
	"fmt"
 	"encoding/binary"
)

func prtb(l string, n uint64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	fmt.Printf("%s: ",l)
	for _, x := range b {
		fmt.Printf("%08b ",x)
	}
	fmt.Printf("\n")
}
*/

func sqrt(r uint64) (root uint64, residue uint64) {
	var onebit uint64
	residue = r
	root = 0
	// Position the one-bit picker until just less than the starting value
	for onebit = 1 << 62; onebit > r; onebit >>= 2 {
	}
	for onebit > 0 {
		x := root | onebit // Current root plus onebit
		if residue >= x {  // Room to subtract?
			residue -= x      // Yes - deduct from residue
			root = x + onebit // and step root
		}
		root >>= 1   // Slide evolving root 1 bit down the residue
		onebit >>= 2 // Slide the bitpick 1 bit down the root
	}
	return
}

func main() {
	var i, q, r, v, delta uint64
	for i = 1<<64 - 1; i > 0; i /= 7 { // for several numbers
		q, r = sqrt(i)                   // Get root and residue
		v = q*q + r                      // Recalc original value
		delta = v - i                    // Difference, hopefully 0
		nextsq := int64((q+1)*(q+1) - i) // Diff to next square, hopefully +ve
		println(i, q, r, v, delta, nextsq)
	}
}
