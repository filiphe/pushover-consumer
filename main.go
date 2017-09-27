package main

import (
	"fmt"
)

func main() {
	var c Credentials
	c.GetCredentials()

	var ar AuthResp
	ar.Auth(c.Username, c.Password)
	fmt.Println(ar)

	var d DeviceResp
	d.Register("test-device")
	fmt.Println(d)
}
