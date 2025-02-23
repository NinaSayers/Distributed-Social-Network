package network

import (
	"fmt"
	"log"
	"net"
)

func Broadcast(port int) {
	fmt.Println("Broadcasting")
	strAddr := fmt.Sprintf("0.0.0.0:%d", port)

	addr, err := net.ResolveUDPAddr("udp4", strAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	defer conn.Close()

	for {
		_, rAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// //fmt.Printf("Received %d bytes from %v\n", n, rAddr)

		go func(rAddr net.Addr) {
			connChan := make(chan net.Conn)

			go func() {
				respConn, err := net.Dial("tcp", rAddr.String())
				for err != nil {
					respConn, err = net.Dial("tcp", rAddr.String())
				}
				connChan <- respConn
			}()

			addrs, _ := net.InterfaceAddrs()

			ip, _, _ := net.ParseCIDR(addrs[1].String())

			ip = ip.To4()

			respConn := <-connChan
			respConn.Write([]byte{ip[0], ip[1], ip[2], ip[3]})
		}(rAddr)
	}
}
