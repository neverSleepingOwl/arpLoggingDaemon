package main

import (
	"log/syslog"
	"log"
	"github.com/coreos/go-systemd/daemon"
	"time"
)

func init(){
	logWriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
	if e == nil {
		log.SetOutput(logWriter)
	}
}

func main() {
	daemon.SdNotify(false, "READY=1")
	go func() {
		interval, err := daemon.SdWatchdogEnabled(false)
		if err != nil || interval == 0 {
			return
		}
		for {
			daemon.SdNotify(false, "WATCHDOG=1")
			time.Sleep(interval / 3)
		}
	}()
	startLoggingArp()
}
