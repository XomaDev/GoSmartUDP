package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	const port = 8080
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+fmt.Sprint(port))
	if err != nil {
		fmt.Println("Error resolve address:", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listen UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Listening UDP on port:", port)

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error receiving packet:", err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr.String())

		// Echo em back
		_, err = conn.WriteToUDP(buffer[:n], addr)
		if err != nil {
			fmt.Println("Could not echo:", err)
			continue
		}
	}
}
