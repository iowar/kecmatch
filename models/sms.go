package models

import (
	"log"
	"sync"
	"time"
)

type SolidityMethodStatus struct {
	start      time.Time
	goroutines int
	found      bool
	data       string
	iteration  int64
	signature  string
	end        time.Time
	sync.RWMutex
}

func NewSolidityMethodStatus(goroutines int) *SolidityMethodStatus {
	return &SolidityMethodStatus{
		start:      time.Now(),
		goroutines: goroutines,
	}
}

func (s *SolidityMethodStatus) IsFound() bool {
	s.RLock()
	defer s.RUnlock()

	return s.found
}

func (s *SolidityMethodStatus) SetInfo(data string, iteration int64, signature string) {
	s.Lock()
	defer s.Unlock()

	s.found = true
	s.data = data
	s.iteration = iteration
	s.signature = signature
	s.end = time.Now()
}

func (s *SolidityMethodStatus) GetInfo() (string, int64, string) {
	s.RLock()
	defer s.RUnlock()

	return s.data, s.iteration, s.signature
}

func (s *SolidityMethodStatus) Display() {
	s.RLock()
	defer s.RUnlock()

	if s.found {
		log.Printf(
			"\033[31m SIGNATURE FOUND:\n\n data: %s\n signature: %s\n iteration: %d\n goroutines: %d\n exec time: %s\n\033[0m",
			s.data,
			s.signature,
			s.iteration,
			s.goroutines,
			s.end.Sub(s.start),
		)
	} else {
		log.Printf("\033[31m SIGNATURE NOT FOUND!\033[0m")
	}
}
