package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("TCP server is running on localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// goroutineでノンブロッキング処理
		go func() {
			io.WriteString(conn, "processing...")
			time.Sleep(3 * time.Second)
			io.WriteString(conn, "done")
			conn.Close()
		}()
	}
}
