package bedrock

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	date     string
	time     string
	priority string
	msg      string
	lock     *sync.Mutex
}

func (m *message) printMessage() error {
	now := time.Now()
	if m.date == "" {
		m.date = now.Format("2006-01-02")
	}
	if m.time == "" {
		m.time = now.Format("15:04:05")
	}
	if m.priority == "" {
		m.priority = "UNK"
	}
	//fmt.Printf("[%s %s %s] %s", msg.date, msg.time, msg.priority, msg.msg)
	print := fmt.Sprintf("[%s %s %s] %s", m.date, m.time, m.priority, m.msg)
	m.lock.Lock()
	fmt.Println(print)
	m.lock.Unlock()
	return nil
}
