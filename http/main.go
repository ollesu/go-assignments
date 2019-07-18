package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Read func pushes raw body response into this byte slice
	// 1) byteSlice := make([]byte, 99999)
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	lw := logWriter{}
	// writer interface, reader interface
	io.Copy(lw, resp.Body)
}

// logWriter now implements Writer interface!
func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))

	fmt.Println("Just wrote this many bytes:", len(byteSlice))
	return len(byteSlice), nil
}

// Reader interface handles all different sources of data (http request, text file, user input from cli, coming into the programme
// it translates it into the common output data - []byte

// every source of data (eg http request body, user input) implements a Reader
