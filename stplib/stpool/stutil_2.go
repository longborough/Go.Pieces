// ST Pool analysis
package main

import (
	"fmt"
	"github.com/longborough/stplib"
)
       
func main() {
	var nP float64 = 15700800         // Pool size
	var n1 int64   = 35               // ST Pool N1 (cycles to timeout)
//	var d  float64 = 8737             // Dispense rate per second
	var rP float64 = 0.1              // Proportion that get released
	var tR float64 = 29.0             // Mean time to release (minutes)
	var tI float64 = 0.98             // Target traffic intensity
	fmt.Printf("n1 tR rtH r0 r1 r2 r3 nU occ pU nC tC cC\n")	
	for tR = 5 ; tR <= 200 ; tR += 5 {
		for n1 = 2 ; n1 <= 100 ; n1++ {
			stats := stplib.STpoolan(nP, n1, nP*tI/tR, rP, tR)
			//		fmt.Printf("\n")
			fmt.Printf("%3d %0.2f %s",n1,tR,stats)
		}
		fmt.Printf("\n")
	}
}      
