// ST Pool analysis
package main

import (
	"fmt"
	"github.com/longborough/stplib"
)
       
func main() {
	printed := false
	model := stplib.STinfo {
		Np: 15700800 ,       // Pool size
		N1: 35       ,       // ST Pool N1 (cycles to timeout)
		Pr: 0.8      ,       // Proportion that get released
		Tr: 15.0     ,       // Mean time to release (minutes)
		Rd: 60*8737  ,       // Mean dispense rate (per minute)
	}
	fmt.Printf("%s\n",stplib.DataHeader)

	for model.Pr = 1.0 ; model.Pr >= 0.099200 ; model.Pr -= 0.01 {
		for model.N1 = 2 ; model.N1 <= 100 ; model.N1++ {
			if stats, err := stplib.STpoolan(model) ; err == nil {
				fmt.Printf("%s\n",stats)
				printed = true
				model.Nu = stats.Nu
			} else { 
				break 
			}
		}
		if printed {
			fmt.Printf("\n")
			printed = false
		}
	}
}      
