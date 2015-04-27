package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// For timestamp
	t := time.Now().Local()

	// Create timestamped file and open it.
	f, err := os.OpenFile("./follogme_"+t.Format("20060102150405")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()

	// Set the log output to stdout and to a file
	log.SetOutput(io.MultiWriter(os.Stdout, f))

	// Create reader buffer
	buffer := bufio.NewReader(os.Stdin)
	// Read each newline
	for {
		fmt.Printf("[i] Log ~: ")
		input, err := buffer.ReadString('\n')
		check(err)
		// Output to stdout and file
		log.Printf(input)
	}
}
