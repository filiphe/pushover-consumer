package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	messagesUrl = "https://api.pushover.net/1/messages.json"
)

type MessageResp struct {
	Messages []Message `json:"messages"`
	User     User      `json:"user"`
	Device   Device    `json:"device"`
	Status   int       `json:"status"`
	Request  string    `json:"request"`
}

func (mr MessageResp) String() string {
	return fmt.Sprintf("Messages: %v\nUser: %v\nDevice: %v\nStatus: %d\nRequest: %s\n", mr.Messages, mr.User, mr.Device, mr.Status, mr.Request)
}

func (mr *MessageResp) LookUpMessages() *MessageResp {
	queryURL := fmt.Sprintf("%s?secret=%s&device_id=%s", messagesUrl, GetSecret(), GetDeviceId())

	resp, err := http.Get(queryURL)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(result, mr)
	if err != nil {
		log.Fatal(err)
	}

	return mr

}

func (mr *MessageResp) GetMessages() []Message {
	return mr.Messages
}
