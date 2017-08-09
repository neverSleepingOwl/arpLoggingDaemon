package main

import (
	"net"
	"strconv"
	"syscall"
	"log"
)


//function, responsible for opening domain socket and accepting connection,
//must be called as go routine because Listen and Accept operations are blocking
func openDomainSocket(chanel chan uint32) {
	//remove socket file if it exists, to prevent errors on restarting after failure
	syscall.Unlink("/tmp/foo")

	l, err := net.Listen("unix", "/tmp/foo")
	if err != nil {
		log.Println("Failed to bind unix domain socket", err.Error())
	}

	log.Println("Bind domain socket bind success, waiting for connection")

	defer l.Close() // close on exit (cleanup)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Println("Failed to accept connection", err.Error())
			return
		}

		log.Println("Domain socket connection accepted")
		readDomainSocket(fd, chanel)
	}
}

//function, receives commands from unix domain socket
//parses commands and sends time to pause logging in specified channel
func readDomainSocket(c net.Conn, chanel chan uint32) {
	buff := make([]byte, 1024) // buffer fo incoming data
	for {
		// commands with length more than 10 chars are ignored
		if size, err := c.Read(buff); err == nil && size < 10 {

			if command, result := parseCommand(buff[:size]); result {

				log.Println("Got command through unix socket: ", command)
				chanel <- command

			}

		} else if err != nil {

			log.Println("Error while reading from unix domain socket", err.Error())
			return

		} else {
		}
	}
}

//function, parses incoming commands, suitable commands are:
//time of pause in seconds or "continue" command which may be '/n' terminated or not
//returns time of delay and bool value of success of operation
func parseCommand(buffer []byte) (uint32, bool) { //TODO add ability to parse '/n' terminated strings and not '/n'
	if i, err := strconv.Atoi(string(buffer[:len(buffer)-1])); err == nil {
		return uint32(i), true //	parsing int
	} else {

		if string(buffer) == "continue\n" || string(buffer) == "continue" {
			return uint32(0), true //	parsing continue command
		}

		return 0, false //	operation failure, command will be ignored
	}
}
