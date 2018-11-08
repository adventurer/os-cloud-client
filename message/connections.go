package message

import (
	"log"
	"os-cloud-client/configure"
	"os-cloud-client/connection"
	"time"

	"github.com/shirou/gopsutil/net"
)

type connectionCnt struct {
	ESTABLISHED int
	TIME_WAIT   int
	CLOSE       int
	LISTEN      int
	OTHER       int
}

func ConnectionStat() {
	var constat connectionCnt
	var now = time.Now()
	connections, err := net.Connections("tcp")
	if err != nil {
		log.Printf("connectionStat err for:%s\n", err)
	}
	for _, v := range connections {
		// log.Printf("%#v\n", v)
		switch v.Status {
		case "ESTABLISHED":
			constat.ESTABLISHED++
		case "TIME_WAIT":
			constat.TIME_WAIT++
		case "CLOSE":
			constat.CLOSE++
		case "LISTEN":
			constat.LISTEN++
		default:
			constat.OTHER++
		}
	}
	var msg connection.Msg
	msg = connection.Msg{Val: constat.ESTABLISHED, Key: "ESTABLISHED", Type: "CONNECTION", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg
	msg = connection.Msg{Val: constat.TIME_WAIT, Key: "TIME_WAIT", Type: "CONNECTION", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg
	msg = connection.Msg{Val: constat.CLOSE, Key: "CLOSE", Type: "CONNECTION", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg
	msg = connection.Msg{Val: constat.LISTEN, Key: "LISTEN", Type: "CONNECTION", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg
	msg = connection.Msg{Val: constat.OTHER, Key: "OTHER", Type: "CONNECTION", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg
	msg = connection.Msg{Val: constat.OTHER + constat.LISTEN + constat.CLOSE + constat.TIME_WAIT + constat.ESTABLISHED, Type: "CONNECTION", Key: "TOTAL", Created: now, Uid: configure.AppConf.Key}
	Msg <- msg

}
