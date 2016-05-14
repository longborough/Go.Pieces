// Extension Tube Calculations
// Arguments: <focal length> <shortest focus> <tube length...>
package main

import "fmt"
import "flag"
import "strconv"

func main() {
	doloop()
}
func doloop() {
	var f0 float64                // Nominal lens focal length
	var fm float64                // Focal length at minimum focus distance
	var md float64                // Minimum focus distance w/out extension
	var el float64				  // Extension tube length
	var xd float64                // Minimum focus with extension
	var yd float64                // Maximum focus with extension
	var err error				  // Error return

	flag.Parse()
	if flag.NArg() < 3 {
		println("Sorry, need 3+ args")
		return
	}
	f0, err = strconv.ParseFloat(flag.Arg(0),64)
	if err != nil {
		println("Sorry - bad focal length")
		return
	}
	md, err = strconv.ParseFloat(flag.Arg(1),64)
	if err != nil {
		println("Sorry - bad minimum distance")
		return
	}
	mdm := 1000 * md
	fm = f0*mdm/(f0+mdm)
	fmt.Printf("Lens: %5.0f mm, Min focus: %0.3f m, Fmin %0.1f mm\n", f0, md, fm)
	for i := 2; i < flag.NArg(); i++ {
		el, err = strconv.ParseFloat(flag.Arg(i),64)
		if err == nil {
			xd = f0*(f0+el)/el
			yd = fm * (f0 + el) / (f0 + el - fm)
			fmt.Printf("Ex: %7.0f     %0.1f -- %0.1f\n", el, xd, yd)
		} else {
			fmt.Printf("Sorry, can't evaluate %s\n",flag.Arg(i))
		}
	}
}

