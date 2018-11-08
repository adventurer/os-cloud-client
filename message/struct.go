package message

import "os-cloud-client/connection"

var Msg = make(chan connection.Msg)
