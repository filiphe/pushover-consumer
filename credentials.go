package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Credentials holds username and password for authenticating client towards Pushover
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetCredentials reads 'credentials.json' and loads a Credentials struct with the information
func (c *Credentials) GetCredentials() *Credentials {
	jf, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Couldn't read 'credentials.json': %v", err)
	}
	err = json.Unmarshal(jf, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
