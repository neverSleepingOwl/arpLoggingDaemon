package main

//This file contains core function and function to implement pause.

import (
	"arpLoggingDaemon/src/CustomTimer"
	"github.com/google/gopacket/layers"
	"log"
)

var (
	logEnabled bool   = true
	arpMissed  uint32 = 0
)

//core function, runs all goroutines, holds logic of program
//in main loop check whether arp request is received and logging is enabled
// if so, writes data to system log
func startLoggingArp() {
	rawChanel := make(chan layers.ARP)
	domainChanel := make(chan uint32)

	go openDomainSocket(domainChanel)
	go receiveArp(rawChanel)
	go pause(domainChanel)

	for {
		if message, ok := ParseArp(<-rawChanel); ok {

			if logEnabled {
				log.Println(message) //TODO add system log
			} else {
				arpMissed++
			}

		}
	}
}

func pause(chanel chan uint32) {
	timer := CustomTimer.Init()
	for {

		if delay := <-chanel; delay > 0 && logEnabled{
			logEnabled = false

			timer.Set(delay)
			timer.Run()

			go func() {
				<-timer.Expired
				logEnabled = true
				log.Println("Requests missed: ", arpMissed)
			}()
		}else if delay > 0{
			timer.Add(delay)
		}else{
			timer.Stop()
		}

	}
}
