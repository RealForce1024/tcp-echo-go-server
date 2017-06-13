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
	CONN_SERVER_HOST = "localhost"
	CONN_SERVER_PORT = ":3333"
	CONN_SERVER_TYPE = "tcp"
)

// 单客户端模拟

//10000次,总共用时0.539791286
//QPS:=>18525.6788306138

//100000次,总共用时5.749476846
//QPS:=>17392.88681014022

func main() {
	conn, err := net.Dial(CONN_SERVER_TYPE, CONN_SERVER_HOST+CONN_SERVER_PORT)
	if err != nil {
		fmt.Println("connect error:=>", err)
		return
	}
	start := time.Now()

	defer conn.Close()

	//userInput := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(conn)

	var n = 100000
	for i := 1; i <= n; i++ {
		_, err := conn.Write([]byte("hello from client=>" + strconv.Itoa(i) + "\n"))
		if err != nil {
			fmt.Println("error", err)
		}

		serverLine, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Println(string(serverLine))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("server error:=>", err)
			os.Exit(2)

		}
	}
	end := time.Since(start).Seconds()

	fmt.Printf("%v次,总共用时%v\n", n, end)
	fmt.Printf("QPS:=>%v\n", float64(n)/end)
}
