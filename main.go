package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		fmt.Errorf("Error when starting server: %v\n", err)
	}
	defer listen.Close()
	
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Errorf("Error when listening: %v\n", err)
		}
		go func(c net.Conn) {
			//var content []byte
			content := make([]byte, 2048)
			_, err := c.Read(content)
			if err != nil {
				fmt.Errorf("Error when reading content: %v\n", err)
			}
			result := strings.Trim(string(content), " ")
			fmt.Println(result)
			fmt.Println(len(result))
		}(conn)
	}
}
