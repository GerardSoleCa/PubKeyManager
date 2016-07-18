package main

import (
	"github.com/GerardSoleCa/PubKeyManager/server"
	"github.com/op/go-logging"
)

var glog = logging.MustGetLogger("database")

func main(){
	glog.Info("Starting PubKeyManager")
	server.Load()
}