package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	host = flag.String("host", "localhost:8080", "the host address to bind to")

	helloWorldString = "Hello, World!"
	helloWorldBytes  = []byte(helloWorldString)
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write(helloWorldBytes)
	})

	fmt.Println("hosting echo service at %v", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		fmt.Printf("fatal: %v\n", err)
	}
}
