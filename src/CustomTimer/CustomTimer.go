package CustomTimer

import (
	"sync"
	"time"
)

//class representing simple timer
//all time values are given in seconds
//if timer expires it's run func writes to Expired chanel
type CTimer struct {
	Expired   chan bool //	chanel to know whether timer is expired or not
	startTime uint32
	running   bool
	mu        sync.Mutex
	delay     uint32
}

//	fill timer with initial values
func Init() (timer CTimer) {
	return CTimer{make(chan bool), 0, false, sync.Mutex{}, 0}
}

// set timer's delay if timer isn't running
func (t *CTimer) Set(delay uint32) {
	if !t.running {
		t.delay = delay //	no need in mutex.Lock because we know that timer isn't running
	}
}

// stop timer
func (t *CTimer) Stop() {
	if t.running {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.delay = 0
		t.Expired <- true
	}
}

//add some time to timers delay (if timer is set to 10 and before it expires we'll add 6 seconds
// it will expire by 16 seconds from the beginning)
func (t *CTimer) Add(delay uint32) {
	if t.expiresFromNow() > 0 {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.delay += delay
	} else {
		t.Set(delay)
	}
}

//returns remaining time in seconds
func (t *CTimer) expiresFromNow() uint32 {
	if exp := t.startTime + t.delay - uint32(time.Now().Unix()); exp > 0 {
		return exp
	} else {
		return 0
	}
}

//start timer, write to channel if it expires
//calls go routine, so we us mutexes when read/write to delay field to prevent race
func (t *CTimer) Run() {
	if !t.running {
		t.running = true
		t.startTime = uint32(time.Now().Unix())
		go func() {
			for {
				if t.expiresFromNow() <= 0 {
					t.Expired <- true
					t.running = false
					return
				}
			}
		}()
	}
}
