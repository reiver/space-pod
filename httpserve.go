package main

import (
	"net/http"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-pod/srv/http"
	"github.com/reiver/space-pod/srv/log"
	_ "github.com/reiver/space-pod/www"
)

func httpserve(tcpaddr string) error {
	log := logsrv.Prefix("httpserve").Begin()
	defer log.End()

	log.Informf("serving HTTP on TCP address: %q", tcpaddr)

	err := http.ListenAndServe(tcpaddr, &httpsrv.Mux)
	if nil != err {
		err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
		log.Errorf("ERROR: %s", err)
		return err
	}

	return nil
}
