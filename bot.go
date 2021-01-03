package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type MessageKind int

const (
	PRIVMSG MessageKind = iota
	NOTICE
)

type Message struct {
	Kind     string
	Sender   string
	Receiver string
	Content  []string
}

func NewMessage(m string) *Message {

	mParts := strings.Split(m, " ")
	return &Message{
		Kind:     mParts[1],
		Sender:   mParts[0],
		Receiver: mParts[3],
		Content:  mParts[4:],
	}
}

func main() {
	conn, err := net.Dial("tcp", "irc.freenode.net:6667")
	if err != nil {
		log.Fatal("Unable to dial freenode!")
	}
	defer conn.Close()

	// RFC states 512 limit for messages, 512 including CRLF
	// buf := make([]byte, 510)
	scanner := bufio.NewScanner(conn)

	// Main read loop
	for scanner.Scan() {
		msg := NewMessage(scanner.Text())
		fmt.Println(msg.Content)
	}

}
