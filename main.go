package main

import (
	"github.com/reiver/space-pod/cfg"
	"github.com/reiver/space-pod/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-pod ⚡")
	blur()

	var httptcpaddr string = cfg.WebServerTCPAddress()
	httpserve(httptcpaddr)
}
