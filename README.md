# ARP logging daemon

Simple systemd Daemon to write ARP requests/replies to syslog. Also you can pause logging by sending commands through 
unix domain socket in format: <integer time in seconds> or word continue. If pause expires daemon will write to syslog 
amount of missed request. Written in pure golang using gopacket lib to capture packets.

## Getting Started

Pull repository, build all *.go files using go build and install in /lib/systemd/system or /usr/lib/systemd/system. Check with the output of the command pkg-config systemd --variable=systemdsystemunitdir to integrate daemon with systemd init system.
If you want to test program without daemonising it in systemd style just replace main.go with your own file main.go, 
and run startLogginArp function from ArpLogger.go.

### Prerequisites
 Required libcap-dev (sudo apt-get install libcap-dev) library. All go packets can be received using "go get" command

```
Give examples
```

### Installing

Pull repository, build all *.go files using go build and install in /lib/systemd/system or /usr/lib/systemd/system. Check with the output of the command pkg-config systemd --variable=systemdsystemunitdir to integrate daemon with systemd init system.

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests
No tests added yet, it was my first learning program in golang.

## Built With
go build 


## License
