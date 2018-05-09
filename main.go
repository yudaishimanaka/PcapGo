package main

import (
	"log"
	"net"
	"os"
	"syscall"
	"fmt"
	"time"
	"flag"
)

const (
	exitCodeOk = iota
	exitCodeErr
)

func htons(host uint16) uint16 {
	return (host&0xff)<<8 | (host >> 8)
}

func main() {
	// get raw sock
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
	if err != nil {
		log.Fatal(err)
	}

	// defer close fd
	defer syscall.Close(fd)

	// get flag
	var (
		d = flag.String("d", "", "-d: device(network interface)")
		w = flag.String("w", "none", "-w: data write pcap file")
		r = flag.String("r", "none", "-r: read pcap file")
	)

	flag.Parse()

	if *d == "" {
		err := "please select device(network interface)."
		log.Fatal(err)
	}

	log.Println(*w, *r)

	// get interface name from argument
	interfaceName := *d

	// check interface exist
	interfaceIndex, err := net.InterfaceByName(interfaceName)
	if err != nil {
		log.Fatal(err)
	}

	// bind interface
	addr := syscall.SockaddrLinklayer{Protocol: htons(syscall.ETH_P_ALL), Ifindex: interfaceIndex.Index}
	if err := syscall.Bind(fd, &addr); err != nil {
		log.Fatal(err)
	}

	// set promiscuous flag
	if err := syscall.SetLsfPromisc(interfaceName, true); err != nil {
		log.Fatal(err)
	}

	file := os.NewFile(uintptr(fd), "")

	// loop analyze raw packet
	for {
		// buffer size is 4096 ~ 65535, AWS spew errors even at 4096 byes
		buffer := make([]byte, 4096)
		now := time.Now()
		num, err := file.Read(buffer)
		if err != nil {
			log.Fatal(err)
			break
		} else {
			binaryData := buffer[:num]
			fmt.Printf("%d:%d:%d.%d\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond())

			err := analyzePacket(binaryData, num)
			if err != nil {
				log.Fatal(err)
				break
			}
		}
	}

	// exit
	os.Exit(exitCodeOk)
}
