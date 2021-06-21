package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func main() {
	go messages()
	Send()
	fmt.Scan()
}

func Send() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/send", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		msg := scanner.Text()

		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}

func messages() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/messages", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg))
	}
}
