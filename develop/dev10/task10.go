package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var timeoutFlag string
	flag.StringVar(&timeoutFlag, "timeout", "10s", "timeout for connection")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]
	timeoutDuration, err := time.ParseDuration(timeoutFlag)
	if err != nil {
		fmt.Printf("Error parsing timeout duration: %v\n", err)
		os.Exit(1)
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeoutDuration)
	if err != nil {
		fmt.Printf("Error connecting to %s:%s: %v\n", host, port, err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	fmt.Printf("Connected to %s:%s\n", host, port)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Connection closed: %v\n", err)
				return
			}
			_, err = os.Stdout.Write(buf[:n])
			if err != nil {
				return
			}
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				return
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				fmt.Printf("Error writing to connection: %v\n", err)
				return
			}
		}
	}()
	chanOs := make(chan os.Signal, 1)
	signal.Notify(chanOs, os.Interrupt, os.Kill)
	<-chanOs
	fmt.Println("Closing connection...")
}
