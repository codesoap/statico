package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var port = flag.Int("p", 8080, "Port to listen on.")

func main() {
	flag.Parse()
	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "Serving files from directory", absPath, "on port", *port)
	handler := logRequest(http.FileServer(http.Dir("")))
	panic(http.ListenAndServe(fmt.Sprint(":", *port), handler))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("15:04:05")
		remoteURL := trimPort(r.RemoteAddr)
		fmt.Printf("%s %s\t%s\t%s\n", now, remoteURL, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func trimPort(addr string) string {
	if endURL := strings.LastIndex(addr, ":"); endURL > 0 {
		return addr[:endURL]
	}
	return addr
}
