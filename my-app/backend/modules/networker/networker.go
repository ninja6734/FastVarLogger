package networker

import (
	"fmt"
	"net"
)

func ReadBuffer(port int, callback func([]byte), testData []string) {
	if testData != nil {
		fmt.Println("Using test data:")
		return
	}
	//magic number: FVLR as bytes so 46 56 4c 52 :   4 bytes
	//data: 2042 bytes
	//end codec: 2 bytes
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error listening on UDP port %d: %v\n", port, err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 2048) // Buffer to hold incoming data

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP: %v\n", err)
			continue
		}

		// Process the received data
		callback(buffer[:n])
	}
}
