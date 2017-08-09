package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"log"
)

func receiveArp(chanel chan layers.ARP) { //TODO add normal exceptions
	Interface, err := net.InterfaceByName("wlp2s0")

	if err != nil {
		log.Println("Error, incorrect inteface: ", err.Error())
		return
	}

	//create pcap handle
	handle, erro := pcap.OpenLive(Interface.Name, 65536, true, pcap.BlockForever)

	if erro != nil {
		log.Println("wrong pcap handle", erro.Error())
		return
	}

	defer handle.Close() // cleanup on exit

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	in := packetSource.Packets()

	//loop to listen for packets
	for {
		var packet gopacket.Packet
		select {
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			chanel <- *arp
		}
	}
}
