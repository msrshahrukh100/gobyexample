package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "myapp.log", "The path of the log file")
	level := flag.String("level", "ERROR", "Log level to search for")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal("Error opening the file ", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}

}
