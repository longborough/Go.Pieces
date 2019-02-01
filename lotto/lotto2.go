// lotto: Generates random entries for Euromillions, Thunderball, and Lotto
package main

import "fmt"
// import "math/rand"
import "sort"
import "time"
import "github.com/MichaelTJones/pcg"

type LPcg struct {
	p *pcg.PCG32 
}

// Perm1: generates a random permutation of numbers in [1:n]  

func (r *LPcg) Perm1(n int) []int {
	m := make([]int, n)
	m[0] = 1
	for i := 1; i < n; i++ {
		j := int(r.p.Bounded(uint32(i + 1)))
		m[i] = m[j]
		m[j] = i + 1
	}
	return m
} 

// Pick1: generates a sorted random permutation of q integers from [1,n]
func (r *LPcg) Pick1(q, n int) []int {
	if q > n {
		q = n
	}

	result := sort.IntSlice((r.Perm1(n))[:q])
	result.Sort()
	return result
}

func main() {
	rng := LPcg {p: pcg.NewPCG32().Seed(1,1).Advance(uint64(time.Now().UnixNano()))}
	fmt.Printf("Euromillions: %v\t%v\n", rng.Pick1(5, 50), rng.Pick1(2, 12))
	fmt.Printf("Thunderball:  %v\t%v\n", rng.Pick1(5, 39), rng.Pick1(1, 14))
	fmt.Printf("Lotto:        %v\n", rng.Pick1(6, 59))
	fmt.Printf("Five:         %v\n", rng.Perm1(5))
}
