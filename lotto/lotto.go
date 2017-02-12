// lotto: Generates random entries for Euromillions, Thunderball, and Lotto
package main

import "fmt"
import "math/rand"
import "sort"
import "time"

// pick: generates a random permutation of q integers from [1,n]
func pick(q, n int) []int {
	if q > n {
		q = n
	}
	result := sort.IntSlice((rand.Perm(n))[0:q])
	for i, _ := range result {
		result[i]++
	}
	result.Sort()
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Euromillions: %v\t%v\n", pick(5, 50), pick(2, 11))
	fmt.Printf("Thunderball:  %v\t%v\n", pick(5, 39), pick(1, 14))
	fmt.Printf("Lotto:        %v\n", pick(6, 59))
	fmt.Printf("Five:         %v\n", rand.Perm(5))
}

/*
func main() {
	rand.Seed(time.Now().UnixNano())
	var list [50]int
	for i := 0; i<1000000; i++ {
		s := pick(5,50)
		for _, r := range s {
			list[r-1]++
		}
	}
	for _, r := range list {
		fmt.Printf("%d\n",r)
	}
}
*/
