package main

import (
	"log"
	"os"
)

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(data)
}
