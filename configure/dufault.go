package configure

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type Appconfig struct {
	Server string
	Key    string
	Port   string
}

var Server *string = flag.String("server", "empty", "Use -server <server ip address>")
var Key *string = flag.String("key", "empty", "Use -key <client key>")
var Port *string = flag.String("port", "empty", "Use -port <client port>")
var AppConf = new(Appconfig)

func init() {
	flag.Parse()
	log.Println(*Server, *Key, *Port)
	if *Server == "empty" || *Key == "empty" || *Port == "empty" {
		data, err := ioutil.ReadFile("./config.json")
		if err != nil {
			log.Fatal("读取配置文件出错：", err)
			return
		}
		err = json.Unmarshal(data, &AppConf)
		if err != nil {
			log.Fatal("解析配置文件出错：", err)
			return
		}
		log.Println("读取了配置文件")
		return
	}

	AppConf.Server = *Server
	AppConf.Key = *Key
	AppConf.Port = *Port
	AppConfJson, _ := json.Marshal(AppConf)
	file, err := os.OpenFile("./config.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("创建配置文件出错")
		return
	}
	_, err = file.WriteString(string(AppConfJson))
	if err != nil {
		log.Fatal("写入配置文件出错")
		return
	}

}
