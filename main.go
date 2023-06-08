package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

var (
	proxyAddr = ""
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	proxyAddr = os.Getenv("PROXY_URL")
	
	fmt.Println("start")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		// fmt.Println(&conn)
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}

func handleConn(dst net.Conn) {
	defer dst.Close()
	// fmt.Println("start2")
	src, err := net.Dial("tcp", proxyAddr)
	// fmt.Println(src)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	go func() {
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}
