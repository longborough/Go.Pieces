// rwalk: Generates a random walk
package main

import "fmt"
import "time"
import "github.com/MichaelTJones/pcg"

type LPcg struct {
	*pcg.PCG32 
}

// Perm0: generates a random permutation of numbers in [1:n]  

func (r *LPcg) Perm0(n int) []int {
	m := make([]int, n)
	m[0] = 0
	for i := 1; i < n; i++ {
		j := int(r.Bounded(uint32(i)))
		m[i] = m[j]
		m[j] = i 
	}
	return m
} 

func main() {
	rng := LPcg {pcg.NewPCG32().Seed(1,1).Advance(uint64(time.Now().UnixNano()))}
//	var limit int = 100
	var x int = 0
	var y int = 0
	for i := 0 ; i < 10000 ; i ++ {
		fmt.Printf("%d,%d\n",x,y)
		inc := int(rng.Bounded(8))
		switch inc {
		case 0: 
			x++
		case 1: y++
		case 2: x--
		case 3: y--
		case 4: 
			x++ 
			y++ 
		case 5: 
			x++ 
			y--
		case 6: 
			x--
			y++
		case 7: 
			x--
			y--
		}
	}
	
}
