package verboten

import (
	"io"
	"net/http"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-pod/srv/http"
)

const path string = "/"

func init() {
	var handler http.Handler = http.HandlerFunc(serveHTTP)

	err := httpsrv.Mux.HandlePath(handler, path)
	if nil != err {
		e := erorr.Errorf("problem registering http-handler with path-mux for path %q: %w", path, err)
		panic(e)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {

	if nil == responsewriter {
		return
	}

	io.WriteString(responsewriter, "bing bong bang it's boomerang")
}
