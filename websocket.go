package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
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
		msg := make([]byte, 512)
		_, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		m := string(msg)
		fmt.Println(m)
		switch m {
		case strings.HasPrefix(m, "!"):
			Sync()
		case strings.HasPrefix(m, "#"):
			continue
		case strings.HasPrefix(m, "R"):
			ReConnect()
		default:
			return
		}
	}

}
