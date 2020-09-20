package main

import (
	"fmt"
	"tcpserver/comm/mylog"
	"tcpserver/server"
)

func main() {
	mylog.Init("./log/s.log", "debug", 21600)
	s := server.New("localhost:9999")
	s.OnNewClient(func(c *server.Client) {
		fmt.Println("new client connected.")
		mylog.LogDebug("New client connected: %s", c.Conn().RemoteAddr().String())
	})
	s.OnNewMessage(func(c *server.Client, m string) {
		fmt.Printf("recieve new message: %s", m)
	})
	s.Listen()
}
