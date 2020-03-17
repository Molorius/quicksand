package bedrock

import (
	"bufio"
	"fmt"
	"strings"
)

type handler struct {
	b *Bedrock
}

func (h *handler) parseOutput(msg string) (message, error) {
	final := message{}
	_, err := fmt.Sscanf(msg, "%s %s %s", &final.date, &final.time, &final.priority)
	if err != nil {
		return final, fmt.Errorf("Could not parse server message: %s", msg)
	}

	i := strings.Index(msg, "]")
	if i < 0 {
		final.date = ""
		final.time = ""
		final.priority = ""
		return final, fmt.Errorf("Could not read message: %s", msg)
	}
	final.date = final.date[1:]                             // remove first bracket
	final.priority = final.priority[:len(final.priority)-1] // remove last bracket
	final.msg = msg[i+2:]
	//fmt.Printf("%d server: %s\r\n", n, value)
	return final, nil
}

// \[([\w\-]+|\-+) ([\w\:]+) (\w+)\] (.*)
func (h *handler) Write(p []byte) (n int, e error) {
	reader := strings.NewReader(string(p))
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		msg := scanner.Text()
		m, err := h.parseOutput(msg)
		if err != nil {
			m.priority = "ERROR"
			m.msg = "Unknown value: " + msg
		} else {

		}
		m.printMessage()
	}
	return len(p), nil
}
