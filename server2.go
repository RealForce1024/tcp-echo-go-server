// Copyright 2015 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2015.php
//
// This code may not conform to popular Go coding idioms

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf("Error: %s : %s\n", msg, err.Error())
		os.Exit(1)
	}
}

func readln(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = input[:len(input)-1]
	if input[len(input)-1] == '\r' {
		input = input[:len(input)-1]
	}
	return input, nil
}

func main() {
	port := flag.String("p", "", "port")
	flag.Usage = func() {
		fmt.Printf("Syntax:\n\techoserver [flags]\nwhere flags are:\n")
		flag.PrintDefaults()
	}

	fmt.Printf("echoserver v 0.90 by Jim Lawless\n")
	flag.Parse()

	if *port == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Printf("\n\nPort: %s\n", *port)
	l, err := net.Listen("tcp", ":"+*port)
	handleError(err, "Listen()")

	defer l.Close()
	fmt.Println("Listening...")
	conn, err := l.Accept()
	handleError(err, "Accept()")

	defer conn.Close()
	fmt.Printf("Connected to : %v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		_, err = fmt.Fprintf(conn, "Enter the word 'quit' (with no quotes) to exit.\r\n")
		handleError(err, "first Fprintf()")

		str, err := readln(reader)
		handleError(err, "readln()")
		if str == "quit" {
			fmt.Println("Quitting.")
			os.Exit(0)
		}
		fmt.Println("Input:" + str)
		_, err = fmt.Fprintf(conn, "You said: %s\r\n", str)
		handleError(err, "second Fprintf()")
	}
}