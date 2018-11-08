package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os-cloud-client/configure"
	"strings"
	"sync"
	"time"
)

var (
	Conn      net.Conn
	IsConnect bool
)

func init() {
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go Connect(wg)
	wg.Wait()
	go keepAlive()
}

func keepAlive() {
	for {
		_, err := Conn.Write([]byte("ping\n"))
		if err != nil {
			log.Println("err on connection,wait for 10 Sec try agin:", err.Error())
			time.Sleep(10 * time.Second)
			Conn, err = net.Dial("tcp", configure.AppConf.Server+":"+configure.AppConf.Port)
			if err != nil {
				log.Println("Error reconnecting:", err)
				os.Exit(1)
			}
			IsConnect = true
			log.Println("reconnected...")
		}
		// 10 seconds a heart beat package
		time.Sleep(10 * time.Second)
	}
}

func Connect(wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	Conn, err = net.Dial("tcp", configure.AppConf.Server+":"+configure.AppConf.Port)
	if err != nil {
		log.Println("Error connecting:", err)
		os.Exit(1)
	}
	IsConnect = true
}

func HandleWrite(wg *sync.WaitGroup, msg Msg) {
	defer wg.Done()
	var err error
	// 序列化数据
	b, _ := json.Marshal(msg)
	writer := bufio.NewWriter(Conn)
	_, err = writer.Write(b)
	//_, e := conn.Write(b)
	if err != nil {
		fmt.Println("Error to send message because of ", err.Error())
		os.Exit(-1)
	}

	// 增加换行符导致server端可以readline
	//conn.Write([]byte("\n"))
	writer.Write([]byte("\n"))
	writer.Flush()
}

func HandleRead(wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(Conn)
	message, err := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	if err != nil {
		IsConnect = false
		log.Print("Error to read message because of:", err)
	}
	// log.Printf("receve:%#v\n", message)
}
