package utils

import (
	"fmt"
	"net"
	"os"
)

var (
	IP string
)

func SetupIp() {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			IP = ipv4.String()
			fmt.Println("IPv4: ", IP)
		}
	}
}
