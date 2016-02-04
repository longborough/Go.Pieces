package main

import "github.com/longborough/sesshash"
import "fmt"
import "flag"
import s "strings"

// import "time"

func doargs() {
	flag.Parse()
	var i int
	var crn string
	for i = 0; i < flag.NArg(); i++ {
		crn = s.ToUpper(flag.Arg(i))
		fmt.Printf("%6d  %s\n", sesshash.SessOrd(crn, 19997), crn)
	}
}
func main() {
	doloop()
}
func doloop() {
	ekey := make([]int, 19997, 19997)
	var i int
	var crn string
	var hash int
	//	t := time.Now()
	for i = 0; i < 100000; i++ {
		crn = fmt.Sprintf("DXA%05d", i)
		hash = sesshash.SessOrd(crn, 19997)
		ekey[hash] += 1
		if ekey[hash] > 9 {
			fmt.Println(crn, ekey[hash], hash)
		}
	}
	//	fmt.Println("Ord,Num")
	//	for i = 0; i < len(ekey); i++ {
	//		fmt.Printf("%d,%d\n",i,ekey[i])
	//	}
}
