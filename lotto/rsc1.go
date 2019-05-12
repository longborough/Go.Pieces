// rwalk: Generates a random walk
package main

import "fmt"
import "time"
import "math"
import "math/rand"

func main() {
	rng := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	var limit int = 1000
	for i := 0 ; i < limit ; i ++ {
		r := 10.0
		phis, phic := math.Sincos(rng.Float64() * 1 * math.Pi)
		thetas, thetac := math.Sincos(rng.Float64() *2 * math.Pi) 
		x:= r * phic * thetac
		y:= r * phic * thetas
		z:= r * phis
		fmt.Println("0,0,0")
		fmt.Printf("%f,%f,%f\n",x,y,z)
	}
	
}
