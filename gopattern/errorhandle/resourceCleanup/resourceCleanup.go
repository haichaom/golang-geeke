package main

import (
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r, err := os.Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer Close(r)
}
