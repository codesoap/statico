package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var port = flag.String("port", "8080", "Port to listen on.")

func main() {
	flag.Parse()

	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "Serving files from directory", absPath, "on port", *port)

	requestLogger := requestLogger{handler: http.FileServer(http.Dir(""))}
	panic(http.ListenAndServe(":" + *port, requestLogger))
}

type requestLogger struct {
	handler http.Handler
}

func (h requestLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(rw, r)
	now := time.Now().Format("15:04:05")
	remoteURL := trimPort(r.RemoteAddr)
	fmt.Printf("%s %s\t%s\t%s\n", now, r.Method, remoteURL, r.RequestURI)
}

func trimPort(addr string) string {
	if endURL := strings.LastIndex(addr, ":"); endURL > 0 {
		return addr[:endURL]
	}
	return addr
}
