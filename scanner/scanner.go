package main

import(
	"fmt"
	"net"
)

func main(){

	_, err := net.Dial("tcp", "scanme.nmap.org:22")
	if err == nil {
		fmt.Println("Connection successful")
	}
}