package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

const (
	wsURL  = "wss://client.pushover.net/push"
	origin = "http://localhost/"
)

func Connect() {
	ws, err := websocket.Dial(wsURL, "", origin)
	defer ws.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ws.Write([]byte(fmt.Sprintf("login:%s:%s\n", GetDeviceId(), GetSecret())))
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := make([]byte, 1)
		_, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: Do some logging instead
		code := msg[0]
		switch {
		case code == 33: // New message received ("!")
			var mr MessageResponse
			mr.LookUpMessages()
			messages := mr.GetMessages()
			message := messages[len(messages)-1]
			fmt.Println(message)
			message.SetAsLatest()
		case code == 35: // Keep-alive message ("#")
			fmt.Println("Keep-alive message received")
		case code == 69: // Error message ("E")
			fmt.Println("Something horrible happened!")
			fmt.Println("Try to login again.")
			return
		case code == 82: // Reconnect message ("R")
			ws.Close()
			// TODO: Close connection and reconnect gracefully
		default:
			fmt.Println("Didn't match anything")
		}
		// TODO: Handle emergency messages as per https://pushover.net/api/client#p2
	}

}
