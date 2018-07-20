package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Device found:")
	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Desc:", device.Description)
		fmt.Println("address:", device.Addresses)
		for _, address := range device.Addresses {
			fmt.Println(address.IP)
			fmt.Println(address.Netmask)
		}
	}
}
