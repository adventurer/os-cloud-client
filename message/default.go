package message

import (
	"time"
)

func init() {
	go runMessage()
}

func runMessage() {
	for {
		ConnectionStat()
		Register()
		time.Sleep(5 * time.Second)
	}
}
