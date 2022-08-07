package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host   = "www.baidu.com"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	conn, err := net.Dial("tcp", remote)

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("Error connect: ", remote)
		return
	}

	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	conn.Close()
}
