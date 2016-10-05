package main

import "network/telnet/server"

func main() {
	server := server.NewServer(":1400")
	server.Run()
}
