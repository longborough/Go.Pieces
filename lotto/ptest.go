// ptest: Testing prime number generation
package main

import "fmt"
import "crypto/rand"

func main() {
	for i := 0; i < 100; i++ {
		p, _ := rand.Prime(rand.Reader,16)
		fmt.Println(p)
	}
}
