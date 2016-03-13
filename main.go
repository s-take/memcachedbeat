package main

import (
	"os"

	"./beater"
	"github.com/elastic/beats/libbeat/beat"
)

var Version = "0.0.1"
var Name = "memcachedbeat"

func main() {
	err := beat.Run(Name, Version, beater.New())
	if err != nil {
		os.Exit(1)
	}
}
