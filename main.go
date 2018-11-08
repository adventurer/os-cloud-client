package main

import (
	"log"
	"os-cloud-client/connection"
	"os-cloud-client/message"
	"sync"
	"time"
)

type Appconfig struct {
	Server string
	Key    string
	Port   string
}

func main() {
	go begin()
	log.Println("server init success...")
	time.Sleep(1 * time.Hour)
}

func begin() {
	var wg = new(sync.WaitGroup)
	for {
		if connection.IsConnect {
			var msg connection.Msg
			msg = <-message.Msg
			wg.Add(2)
			connection.HandleWrite(wg, msg)
			connection.HandleRead(wg)
			wg.Wait()
		}
	}
}
