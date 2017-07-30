// Testing ehllapi
package main

import (
	"fmt"
	"regexp" 
)

func main() {
	keytable := map[string]string {
		"{tab}":   "@T",
		"{enter}": "@E", 
	}  
	keys := make([]string, len(keytable))
	for k, _ := range keytable {
		keys = append(keys, k)
	}
	line := "a{lard}a{tab}b{tab}c{enter}{boggis}x"
	orig := line
	for _, k := range keys {
		line = regexp.MustCompile(k).ReplaceAllLiteralString(line,keytable[k])
		}
	fmt.Printf("%q\n",line)
	fmt.Printf("%q\n",orig)
}

