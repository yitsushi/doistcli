package main

import (
	"log"
)

func main() {
	conf, err := parseConfiguration()

	if err != nil {
		panic(err)
	}

	log.Println(conf)
}
