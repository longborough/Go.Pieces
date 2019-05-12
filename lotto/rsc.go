// rwalk: Generates a random walk
package main

import "fmt"
import "time"
import "math"
import "github.com/MichaelTJones/pcg"

func main() {
	rng := pcg.NewPCG32().Seed(1,1).Advance(uint64(time.Now().UnixNano()))
	var limit int = 10000
	for i := 0 ; i < limit ; i ++ {
//		r := 0.01 * float64(rng.Bounded(1000))

		r := 10.0
		phis, phic := math.Sincos(float64(rng.Bounded(36000)) * math.Pi / 18000)
		thetas, thetac := math.Sincos(float64(rng.Bounded(36000)) * math.Pi / 18000) 
		x:= r * phic * thetac
		y:= r * phic * thetas
		z:= r * phis
		fmt.Printf("%f,%f,%f\n",x,y,z)
	}
	
}
