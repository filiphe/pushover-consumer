package main

//import (
//	"fmt"
//)

// TODO: Refactor the shit out of this
// Fetch more messages on startup before opening websocket connection
// and continuously getting messages.
// TODO: Only register device when necessary.
// TODO: Prompt for credentials and save them in some proper place.
func main() {
	// var c Credentials
	// c.GetCredentials()

	// var ar AuthResp
	// ar.Auth(c.Username, c.Password)
	// fmt.Println(ar)

	// var d DeviceResp
	// d.Register("test-device-laptop")
	// fmt.Println(d)

	// var mr MessageResp
	// mr.LookUpMessages()
	// messages := mr.GetMessages()
	// fmt.Println(messages)
	//length := len(messages)
	//lastMessage := messages[length-1]
	//fmt.Println(lastMessage)
	//lastMessage.SetAsLatest()
	//messages[len(messages)-1].UpdateToLatest()

	Connect()
}
