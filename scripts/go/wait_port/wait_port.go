package main

import (
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port <= 0 {
		os.Exit(1)
	}
	timeout := 30 * time.Second
	if len(os.Args) >= 3 {
		if s, err := strconv.Atoi(os.Args[2]); err == nil && s > 0 {
			timeout = time.Duration(s) * time.Second
		}
	}
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		c, err := net.DialTimeout("tcp", "127.0.0.1:"+strconv.Itoa(port), 2*time.Second)
		if err == nil {
			c.Close()
			os.Exit(0)
		}
		time.Sleep(500 * time.Millisecond)
	}
	os.Exit(1)
}
