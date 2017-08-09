# ARP logging daemon

Simple systemd Daemon to write ARP from interface wlp2s0 requests/replies to syslog. Also you can pause logging by sending commands through 
unix domain socket in format: <integer time in seconds> or word continue. If pause expires daemon will write to syslog 
amount of missed request. Written in pure golang using gopacket lib to capture packets.

## Getting Started
It's important, that daemon must be run with root privileges, or you should run something like :setcap cap_net_raw+ep <your_appp_name>.
Pull repository, build all *.go. Than you can just run program with root privilege. Integrating daemon with systemd is described in install section.

### Prerequisites
 Required libcap-dev library. All go packets can be received using "go get" command.
 Also it's usefull to use socat utility. Example for Debian/Ubuntu:

```
sudo apt-get update
sudo apt-get install libcap-dev
sudo apt-get install socat
go get github.com/google/gopacket
go get github.com/google/gopacket/layers
go get github.com/google/gopacket/pcap
go get github.com/coreos/go-systemd/daemon
sudo socat - UNIX-CONNECT:/tmp/foo   // use this then daemon is already working to pause/continue logging

```
 
### Installing

Pull repository, build all *.go files using go build and install in /usr/bin/arpLoggingDaemon or wherever you want,
if you change ExecStart value in ArpLogginDaemon.service.
You should also locate ArpLogginDaemon.service in /etc/systemd/system/ or write your own systemd script.
After that execute the following commands:

```
systemctl status myunit
systemctl enable myunit // if status: disabled
systemctl start myunit
```
If systemctl prints errors, fix it before start.

## Running the tests
No tests added yet, it was my first learning program in golang.

## Built With
go build *.go should work just fine, also you print all filenames manually or write your own build script.


## License
