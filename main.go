package main

import (
	"flag"
	"fmt"
	"os"
	"golang-rest-api-starter/config"
	"golang-rest-api-starter/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}