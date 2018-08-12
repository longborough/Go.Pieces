// deal: Generates bridge deals
package main

import "fmt"
import "sort"
import "time"
import "github.com/MichaelTJones/pcg"

type LPcg struct {
	*pcg.PCG32 
}

type Card struct {
	serial uint32
	suit string
	rank string
}

type Hand struct {
	serial uint32
	cards []Card
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
	suits := 
	rng := LPcg {pcg.NewPCG32().Seed(1,1).Advance(uint64(time.Now().UnixNano()))}
	deal := rng.Perm0(52)
	hands := make([]sort.IntSlice,4)
	for pl := 0 ; i < 4 ; i ++ {
		hands[pl] = sort.IntSlice(deal[13*pl:13*(pl+1)])
	}
	
}
