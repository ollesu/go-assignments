package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Reader pushes raw data response into this byte slice
	byteSlice := make([]byte, 99999)
	resp.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))
}

// Reader interface handles all different sources of data (http request, text file, user input from cli, coming into the programme
// it translates it into the common output - []byte
