package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
	"github.com/NinaSayers/Distributed-Social-Network/dns/utils"
)

var broadcastPort map[string]int = map[string]int{
	"distnet.com.":     41234,
	"api.distnet.com.": 53123,
}

type handler struct{}

func (h *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		//fmt.Println(domain)
		IP, err := h.broadcast(broadcastPort[domain])
		//fmt.Println("IP Found:", IP)
		if err == nil {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   IP,
			})
		}
	}
	w.WriteMsg(&msg)
}

func (h *handler) broadcast(port int) (net.IP, error) {

	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: port,
	}

	conn, err := net.DialUDP("udp4", nil, &raddr)
	if err != nil {
		return nil, err
	}

	IP, Port, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return nil, err
	}

	_, err = conn.Write([]byte{})
	if err != nil {
		return nil, err
	}
	conn.Close()

	address := fmt.Sprintf("%s:%s", IP, Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	for err != nil {
		return nil, err
	}

	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	tcpConn, err := lis.AcceptTCP()
	if err != nil {
		return nil, err
	}
	//fmt.Println("AcceptTCP Connection")
	ip, err := utils.DeserializeMessage(tcpConn)
	if err != nil {
		return nil, err
	}
	//fmt.Println(ip)

	return ip, nil
}

func main() {
	srv := &dns.Server{Addr: ":53", Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
