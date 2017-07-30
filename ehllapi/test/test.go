// Testing ehllapi
package main

import (
	"fmt"
	eh "github.com/longborough/ehllapi"
)

func main() {
	fmt.Println(eh.GetVersion())
	mysess, err := eh.NewSession("A")
	fmt.Println(mysess)
	fmt.Println(err)
}
