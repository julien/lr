package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"
)

const (
	defaultport = "8000"
)

var (
	valid = regexp.MustCompile(`\d{4}`)
	port  string
	root  string
)

func init() {
	flag.StringVar(&port, "port", defaultport, "default port")
	flag.StringVar(&root, "root", ".", "root directory")
}

func main() {
	flag.Parse()
	fmt.Printf("starting server: 0.0.0.0:%s - root directory: %s\n", port, path.Dir(root))
	http.Handle("/", noIconHandler(fileHandler(root)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func checkPort(port string) string {
	if !valid.MatchString(port) {
		port = defaultport
	}
	return port
}

func fileHandler(root string) http.Handler {
	return http.FileServer(http.Dir(root))
}

func noIconHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := fmt.Sprintf("%s", r.URL)
		if u == "/favicon.ico" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
