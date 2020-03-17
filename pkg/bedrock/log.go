package bedrock

import (
	"fmt"
	"time"
)

type message struct {
	date     string
	time     string
	priority string
	msg      string
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
	fmt.Printf("[%s %s %s] %s\r\n", m.date, m.time, m.priority, m.msg)
	return nil
}
