package details

import (
	"log"
	"net"
	"os"
)

func GetHostName() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}

func GetIP() net.IP {
	log.Println("Test1")
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Test")

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
