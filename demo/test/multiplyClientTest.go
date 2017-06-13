package main

import (
	"net"
	"fmt"
	"time"
	"strconv"
	"io"
	"os"
	"bufio"
)

const (
	_CONN_SERVER_HOST = "localhost"
	_CONN_SERVER_PORT = ":3333"
	_CONN_SERVER_TYPE = "tcp"
)

//模拟多客户端 短连接
// 1k->6.62s
// 3k->21s
// 连接资源有限
func main() {

	start := time.Now()

	for i := 1; i <= 1000; i++ {
		conn, err := net.Dial(_CONN_SERVER_TYPE, _CONN_SERVER_HOST+_CONN_SERVER_PORT)
		response := bufio.NewReader(conn)
		if err != nil {
			fmt.Println("connect error:=>", err)
			return
		}

		//defer conn.Close()

		//for i := 0; i < 50000; i++ {
		_, err = conn.Write([]byte("hello from client=>" + strconv.Itoa(i) + "\n"))

		if err != nil {
			fmt.Println("error", err)
		}
		//}

		serverLine, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Println("server response:=>", string(serverLine))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("server error:=>", err)
			os.Exit(2)

		}
		fmt.Println("客户端=>", i, "用时", time.Since(start).Seconds())
		conn.Close()
	}
	fmt.Println("总时间:", time.Since(start).Seconds())

}
