package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	latestURL = "https://api.pushover.net/1/devices/%s/update_highest_message.json"
)

type Message struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	App      string `json:"app"`
	Aid      int    `json:"aid"`
	Icon     string `json:"icon"`
	Date     int64  `json:"date"`
	Priority int    `json:"priority"`
	Acked    int    `json:"acked"`
	Umid     int    `json:"umid"`
	Title    string `json:"title"`
}

func (m Message) String() string {
	return fmt.Sprintf("Title: %s\nTime: %s\nMessage: %s\nID: %d\n", m.Title, time.Unix(m.Date, 0), m.Message, m.ID)
}
