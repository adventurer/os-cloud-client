package message

import (
	"encoding/json"
	"log"
	"os-cloud-client/configure"
	"os-cloud-client/connection"
	"time"

	"github.com/shirou/gopsutil/host"
)

type Client struct {
	Uid    string
	Host   string
	Os     string
	Uptime uint64
	Ip     string
	Alive  time.Time
}

func Register() {
	info, err := host.Info()
	if err != nil {
		log.Println(err)
		return
	}
	data := Client{Uid: configure.AppConf.Key, Host: info.Hostname, Os: info.OS, Uptime: info.Uptime, Alive: time.Now()}
	infoJson, _ := json.Marshal(data)
	now := time.Now()
	var msg connection.Msg
	msg = connection.Msg{Type: "INFO", Created: now, Data: string(infoJson)}
	Msg <- msg
}
