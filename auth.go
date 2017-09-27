package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AuthResp struct {
	Status  int    `json:"status"`
	ID      string `json:"id"`
	Secret  string `json:"secret"`
	Request string `json:"request"`
}

func (ar AuthResp) String() string {
	return fmt.Sprintf("status: %d\nid: %s\nsecret: %s\nrequest: %s\n", ar.Status, ar.ID, ar.Secret, ar.Request)
}

func (ar *AuthResp) Auth(username, password string) *AuthResp {
	data := url.Values{}
	data.Set("email", username)
	data.Set("password", password)

	resp, err := http.PostForm("https://api.pushover.net/1/users/login.json", data)
	if err != nil {
		log.Fatalf("Unable to login: %+v", err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(result, ar)
	if err != nil {
		log.Fatal(err)
	}

	ar.SaveSecret()
	return ar
}

func (ar *AuthResp) SaveSecret() {
	ioutil.WriteFile("secret.txt", []byte(ar.Secret), 0600)
}

func GetSecret() string {
	secret, err := ioutil.ReadFile("secret.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(secret)
}
