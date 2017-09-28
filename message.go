package main

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"net/url"
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

type MessageResp struct {
	Messages []Message `json:"messages"`
	User     User      `json:"user"`
	Device   Device    `json:"device"`
	Status   int       `json:"status"`
	Request  string    `json:"request"`
}
