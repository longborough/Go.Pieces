// Extension Tube Calculations
// Arguments: <focal length> <shortest focus> <tube length...>
package main

import "fmt"
import "flag"
import "math"
import "strconv"

func main() {
	doloop()
}

func getext() []float64 {
	var list []float64
	list = make([]float64,flag.NArg()-2) 
	for i := 2; i < flag.NArg(); i++ {
		el, err := strconv.ParseFloat(flag.Arg(i),64)
		if err == nil {
			list[i-2] = el
		} else {
			list[i-2] = math.NaN()
		}
	}
	return list
}

func doloop() {
	var f0 float64                // Nominal lens focal length
	var fm float64                // Focal length at minimum focus distance
	var md float64                // Minimum focus distance w/out extension
	var err error				  // Error return
	var exlist []float64

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
	exlist = getext()
	mdm := 1000 * md
	fm = f0*mdm/(f0+mdm)
	fmt.Printf("Lens: %5.0f mm, Min focus: %0.3f m, Fmin %0.1f mm\n", f0, md, fm)
	exlist = getext()
	calcext(exlist,0.0,0,f0,fm)
	return
}
func calcext(list []float64, tex float64, i int, f0 float64, fm float64) {
	println("-->",tex,i)
	var j int
	var k int
	var cex float64
	var elx float64
	for j = i; j < len(list); j++ {
		cex = list[j]
		if math.IsNaN(cex) || cex <= 0.0  {
			continue
		}
		elx = tex + cex 
		xd := f0*(f0+elx)/elx
		yd := fm * (f0 + elx) / (f0 + elx - fm)
		fmt.Printf("Ex: %7.0f     %0.1f -- %0.1f  %d  %d  %0.0f \n", elx, xd, yd, j, i, tex)
		for k = j+1; k < len(list); k++ {
			calcext(list,elx,k,f0,fm)
		}
	}
	println("<--",tex,i)
	return
}

