package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handleRequest)

	log.Printf("telemetry server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println()

	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)

	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	bytes, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Body: <%s>\n", bytes)
}
