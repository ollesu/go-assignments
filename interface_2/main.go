package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error opening the file: ", err)
		os.Exit(1)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}
