package bedrock

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type Bedrock struct {
	ServerDir string
	started   bool
	handler   handler
	clients   map[string]int

	clientLock *sync.Mutex
}

func (b *Bedrock) Start() error {

	b.started = true
	b.handler.b = b // save pointer
	b.clients = make(map[string]int)
	b.clientLock = &sync.Mutex{}

	msg := message{msg: "Starting Quicksand", priority: "INFO"}
	msg.printMessage()

	cmd := exec.Command("./bedrock_server")
	cmd.Dir = b.ServerDir
	cmd.Stdout = &b.handler
	cmd.Stdin = os.Stdin
	//cmd.Stderr = os.Stderr
	//err := cmd.Start()
	err := cmd.Run()
	return err
}

func (b *Bedrock) Stop() error {
	b.started = false
	fmt.Println("Stopping, currently blocking.")

	return nil
}

func (b *Bedrock) addClient(client string) error {
	b.clientLock.Lock()
	defer b.clientLock.Unlock()

	if _, ok := b.clients[client]; ok {
		return fmt.Errorf("Client already in list: %s", client)
	}
	b.clients[client] = 1
	return nil
}

func (b *Bedrock) delClient(client string) error {
	b.clientLock.Lock()
	defer b.clientLock.Unlock()

	if _, ok := b.clients[client]; ok {
		delete(b.clients, client)
		return nil
	}
	return fmt.Errorf("Client not in list: %s", client)
}

func (b *Bedrock) Clients() ([]string, error) {
	b.clientLock.Lock()
	defer b.clientLock.Unlock()

	slice := make([]string, 0, len(b.clients))
	for s := range b.clients {
		slice = append(slice, s)
	}
	return slice, nil
}
