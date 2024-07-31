package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	// bs := make([]byte, 1000000)
	// file.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, file)
}
