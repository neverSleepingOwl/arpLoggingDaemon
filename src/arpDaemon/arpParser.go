package main

import (
	"fmt"
	"github.com/google/gopacket/layers"
	"strconv"
)

// function, converts all parameters of request into
// a human-readable string for writing to log
func ParseArp(data layers.ARP) (result string, done bool) { //TODO done value is useless
	result += "[ARP startLoggingArp ]"

	arr := []string{protocolType(data),
		hardwareLength(data), protocolLength(data),
		operationCode(data), senderMac(data),
		senderIp(data), targetMac(data), targetIp(data)}

	arr2 := []string{" | PROTOCOL TYPE=0x",
		" | HARDWARE LENGTH=", " | PROTOCOL LENGTH= ",
		" | OPCODE=", " | SENDER MAC=", " | SENDER IP=",
		" | TARGET MAC=", " | TARGET IP="}

	for i := range arr {
		result += arr2[i] + arr[i]
	}
	return result + "[ARP]", true
}

//set of functions to convert parameters to strings
//TODO make it simple and clear, this way was used because in prev. version raw ethernet frames were parsed
func protocolType(data layers.ARP) string {
	return fmt.Sprintf("%X", int(data.Protocol))
}

func hardwareLength(data layers.ARP) string {
	return strconv.Itoa(int(data.HwAddressSize))
}

func protocolLength(data layers.ARP) string {
	return strconv.Itoa(int(data.ProtAddressSize))
}

func operationCode(data layers.ARP) string {
	return strconv.Itoa(int(data.Operation))
}

func senderMac(data layers.ARP) string {
	return parseAddr(data.SourceHwAddress)
}

func senderIp(data layers.ARP) string {
	return parseAddr(data.SourceProtAddress)
}

func targetMac(data layers.ARP) string {
	return parseAddr(data.DstHwAddress)
}

func targetIp(data layers.ARP) string {
	return parseAddr(data.DstProtAddress)
}

//function to convert slice of bytes representation of IP/MAC
//into string of byte numbers, split by '.' or ':'
//MAC address contains hex digits, ip  - decimal
func parseAddr(addr []byte) (output string) {
	if len(addr) == 6 {
		for i, element := range addr {
			output += fmt.Sprintf("%X", int(int(element)/256))
			output += fmt.Sprintf("%X", int(int(element)%256))
			if i < 5 {
				output += ":"
			}
		}
	} else {
		for i, element := range addr {
			output += strconv.Itoa(int(element))
			if i < 3 {
				output += "."
			}
		}
	}
	return output
}
