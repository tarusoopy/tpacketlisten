package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	port := flag.String("p", "10000", "Listening Port")
	waittime := flag.Int("w", 1800, "Waiting second until close")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+*port)
	fmt.Println("TCP server is running on localhost:" + *port)
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
			time.Sleep((time.Duration(*waittime)) * time.Second)
			conn.Close()
		}()
	}
}
