// Testing ehllapi
package main

import (
	eh "github.com/longborough/ehllapi"
)

func main() {
	print(eh.GetVersion)
	mysess, err := eh.NewSession("A")
	print(mysess.String())
	print(err)
}
