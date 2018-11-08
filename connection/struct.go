package connection

import (
	"time"
)

type Msg struct {
	Uid     string
	Type    string `json:"type"`
	Data    string `json:"data"`
	Key     string `json:"key"`
	Val     int    `json:"val"`
	Created time.Time
}

type Resp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}
