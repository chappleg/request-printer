package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := ":80"

	if len(os.Args) > 1 {
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid port '%s' was provided: %v", os.Args[1], err)
		}

		port = ":" + strconv.Itoa(portNumber)
	}

	http.HandleFunc("/", handleRequest)

	log.Printf("request printer listening on port %s", port)
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
