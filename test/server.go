package main

import (
	"net"
	"bufio"
	"io"
	"fmt"
	"os"
)

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		fmt.Println("receive:=>", string(line))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("eror", err)
		}
		conn.Write(line)
	}
}

var server = "localhost" // if localhost this can be ignore
var port = ":3545"       // ignore local ip,should add ":"and port

func main() {
	l, err := net.Listen("tcp", port)
	fmt.Println("server listen at", port)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		go echo(conn)
	}

}
