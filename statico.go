package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var port = flag.String("port", "8080", "Port to listen on.")
var path = flag.String("path", ".", "Root directory to serve files from.")

func main() {
	flag.Parse()
	src := ":" + *port

	absPath, _ := filepath.Abs(*path)
	fmt.Fprintln(os.Stderr, "Serving files from directory", absPath, "on", src)

	requestLogger := requestLogger{handler: http.FileServer(http.Dir(*path))}
	panic(http.ListenAndServe(src, requestLogger))
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
