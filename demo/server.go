package main

import (
	"net"
	"bufio"
	"io"
	"fmt"
	"os"
)

const (
	_CONN_HOST = "localhost"
	_CONN_PORT = ":3333"
	_CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(_CONN_TYPE, _CONN_HOST+_CONN_PORT)
	fmt.Println("server listen at", _CONN_PORT)

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		go echo(conn)
	}

}

//使用bufio工具类读取
func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		if err != nil { // 避免client关闭，server端还继续工作
			return
		}

		fmt.Println("receive:=>", string(line))
		conn.Write([]byte("server say:=>" + string(line)))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("eror", err)
		}
		//conn.Write(line)
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
