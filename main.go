package main

import (
	"fmt"
)

func main() {
	//var c Credentials
	//c.GetCredentials()

	//var ar AuthResp
	//ar.Auth(c.Username, c.Password)
	//fmt.Println(ar)

	//var d DeviceResp
	//d.Register("test-device")
	//fmt.Println(d)

	var mr MessageResp
	mr.LookUpMessages()
	messages := mr.GetMessages()
	fmt.Println(messages)
	//messages[len(messages)-1].UpdateToLatest()
}
