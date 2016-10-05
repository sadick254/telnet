package main

import "network/telnetserver"

func main() {
	server := telnetserver.NewServer(":1400")
	server.Run()
}
