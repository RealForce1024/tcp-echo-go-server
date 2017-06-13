package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"io"
)

const (
	CONN_SERVER_HOST = "localhost"
	CONN_SERVER_PORT = ":3333"
	CONN_SERVER_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(CONN_SERVER_TYPE, CONN_SERVER_HOST+CONN_SERVER_PORT)
	if err != nil {
		fmt.Println("connect error:=>", err)
		return
	}

	userInput := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(conn)

	for {
		line, err := userInput.ReadBytes(byte('\n'))
		switch err {
		case nil:
			conn.Write(line)
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("error:=>", err)
			os.Exit(-1)
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
}
