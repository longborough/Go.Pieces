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
	fmt.Printf("Euromillions: %v %v\n",pick(5,50),pick(2,11))
	fmt.Printf("Thunderball:  %v %v\n",pick(5,39),pick(1,14))
	fmt.Printf("Lotto:        %v\n",pick(6,59))
}
