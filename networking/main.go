package main

import (
	"fmt"
	"net"
)

func main() {

}

func listen(port string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	l.Accept()
}
