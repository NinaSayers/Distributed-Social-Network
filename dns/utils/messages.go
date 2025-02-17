package utils

import (
	"encoding/binary"
	"io"
	"net"
	"strconv"
)

func SerializeMessage(ip string, port string) []byte {
	message := []byte(net.ParseIP(ip))
	bs := make([]byte, 4)

	portInt, _ := strconv.Atoi(port)
	binary.LittleEndian.PutUint32(bs, uint32(portInt))
	message = append(message, bs...)

	return message
}

func DeserializeMessage(conn io.Reader) (net.IP, error) {

	ipBuff := make([]byte, 4)
	_, err := conn.Read(ipBuff)
	if err != nil {
		return nil, err
	}

	return net.IP(ipBuff), nil
}
