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

//模拟一个客户端
// 1k->0.07
// 10k->0.77
// 30k->2.33
// 50k->4.06
// 100k->7.61
// 1000k->86.468115049
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

	for i := 1; i <= 1000000; i++ {
		_, err := conn.Write([]byte("hello from client=>" + strconv.Itoa(i) + "\n"))
		if err != nil {
			fmt.Println("error", err)
		}

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
	}
	fmt.Println("总用时:", time.Since(start).Seconds())
}
