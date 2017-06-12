package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var port = "0.0.0.0:9001"

// ->使用bufio向流中读，然后从流中写出
func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("ERROR", err)
		}
		conn.Write(line)
	}
}

func main() {
	// net直接监听端口
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		go echo(conn) //异步并发
		//echo(conn) 阻塞 超过一个客户端服务器端都是阻塞状态
	}
}
