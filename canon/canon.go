package main

import "fmt"
import "net"
import "bufio"

func main() {
	conn, err := net.Dial("tcp", "andygotts.hightor:49152")
	if err != nil {
	println(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	println(status)
	println(err)
}
