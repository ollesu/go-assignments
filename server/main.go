package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// this connects Handler func to incoming URLs whose path begins with /
	http.HandleFunc("/", handler) 
	// starts a server listening to incoming requests
	log.Fatal(http.ListenAndServe("localhost:8000", nil))


	// each request is a struct of type http.Request (contains its URL)
	// upon arrival request is given to handler, which extracts the path /... from the URL
	func handler(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	}
}

// server runs handler for each request so it can serve multiple req. simultaneously 
// --> need to lock a variable for only one use at the same time 

// goroutine - each concurrently executing activity 

// 1. newRequest(type, path)
// 2. if error-statement
// 3. set header for the request
// 4. http.HandlerFunc?? to initialise a handler?
// 5. serveHTTP takes response writer (http.ResponseWriter) and request (*http.Request) as arguments

// Go's HTTP takes an address to listen on and a handler (created by serveHTTP func)

// difference between ListenAndServe and other func 
