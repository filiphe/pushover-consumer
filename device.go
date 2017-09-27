package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type DeviceResp struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Request string `json:"request"`
}

func (d DeviceResp) String() string {
	return fmt.Sprintf("id: %s\nstatus: %d\nrequest: %s\n", d.ID, d.Status, d.Request)
}

func (d *DeviceResp) Register(name string) *DeviceResp {
	data := url.Values{}
	data.Set("secret", GetSecret())
	data.Set("name", name)
	data.Set("os", "O") // os=O is for OpenAPI clients

	resp, err := http.PostForm("https://api.pushover.net/1/devices.json", data)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(result, d)
	if err != nil {
		log.Fatal(err)
	}

	if d.Status == 0 {
		log.Printf("Device \"%s\" is already registered!", name)
	}

	if d.Status == 1 {
		d.SaveDeviceId()
	}

	return d
}

func (d *DeviceResp) SaveDeviceId() {
	ioutil.WriteFile("device_id.txt", []byte(d.ID), 0644)
}

func GetDeviceId() string {
	deviceID, err := ioutil.ReadFile("device_id.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(deviceID)
}
