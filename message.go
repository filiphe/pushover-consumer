package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

type SetLatestMessageResponse struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

func (m Message) String() string {
	return fmt.Sprintf("Title: %s\nTime: %s\nMessage: %s\nID: %d\n", m.Title, time.Unix(m.Date, 0), m.Message, m.ID)
}

func (m *Message) SetAsLatest() {
	data := url.Values{}
	data.Set("secret", GetSecret())
	data.Set("message", strconv.Itoa(m.ID))

	requestURL := fmt.Sprintf(latestURL, GetDeviceId())

	fmt.Println(requestURL)
	fmt.Println(data)
	resp, err := http.PostForm(requestURL, data)
	if err != nil {
		log.Fatalln(resp)
		log.Fatalf("Unable to send request to set latest message: %+v", err)
	}

	response := SetLatestMessageResponse{}

	result, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(result, response)
	// if response.Status != 1 {
	// 	log.Fatalf("Unable to set message as latest")
	// }
	fmt.Println(response)
}
