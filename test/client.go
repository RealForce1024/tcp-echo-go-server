package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"io"
)

var serverPort = ":3545"

func main() {
	conn, err := net.Dial("tcp", serverPort)
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
