package main

import "fmt"
import "flag"
import "strconv"

func main() {
	doloop()
}
func doloop() {
	const optblock = 27998 
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		lrecl, err := strconv.Atoi(flag.Arg(i))
		if err == nil && lrecl > 0 && lrecl <= optblock {
			nrecs := optblock/lrecl
			block := nrecs * lrecl
			fmt.Printf("Lrecl: %5d  Block: %5d  Per block: %5d  Efficiency: %3.1f\n", lrecl, block, nrecs, 100*float32(block)/27998)
		} else {
			fmt.Printf("Sorry, can't evaluate %s\n",flag.Arg(i))
		}
	}
}

