package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	time.Sleep(2*time.Second)
	fmt.Println(time.Since(start).Seconds())
}
