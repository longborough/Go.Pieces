// ST Pool analysis
package main

import (
	"fmt"
	"github.com/longborough/stplib"
)
       
func main() {
	var nP float64 = 15700800         // Pool size
	var n1 int64   = 35               // ST Pool N1 (cycles to timeout)
	var d  float64 = 8737             // Dispense rate per second
	var rP float64 = 0.4              // Proportion that get released
	var tR float64 = 29.0             // Mean time to release (minutes)
	for n1 = 30 ; n1 <= 200 ; n1++ {
		stats := stplib.STpoolan(nP, n1, 60*d, rP, tR)
//		fmt.Printf("\n")
		fmt.Printf("%3d %s",n1,stats)
	}
}      
