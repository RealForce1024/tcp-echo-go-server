package main

import (
	"net"
	"bufio"
	"io"
	"fmt"
	"os"
)

//使用bufio工具类读取
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

// 直接从conn流中读取
func echo2(c net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := c.Read(buf)
		//c.SetDeadline(time.Now()) //可以根据具体情况设置超时时间
		if err != nil {
			//fmt.Println("err:", err)
			return
		}
		data := buf[:n]
		fmt.Println("receive:=>", string(data))
		_, err = c.Write(data)
		switch err {
		case nil:
			//fmt.Println("err:=>", err)
			break
		case io.EOF:
		default:
			fmt.Println("err:=>", err)
		}
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
		//go echo(conn)
		go echo2(conn)
	}

}
