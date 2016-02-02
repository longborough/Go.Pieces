package main

import "github.com/longborough/sesshash"
import "fmt"
import "flag"
import "strconv"
import s "strings"

func doargs() {
	flag.Parse()
	var size int = 19997
	//    var nsize int
	var i int
	var crn string
	for i = 0; i < flag.NArg(); i++ {
		crn = s.ToUpper(flag.Arg(i))
		if nsize, err := strconv.Atoi(crn); err == nil {
			size = nsize
		} else {
			fmt.Printf("%6d,%6d,%s\n", sesshash.SessOrd(crn, size), size, crn)
		}
	}
}
func main() {
	fmt.Printf("Ord,Size,Key\n")
	doargs()
}
