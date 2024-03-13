package main

import (
	"fmt"
	"log"
	"pack-svc/pkg/app"
)

const (
	httpServeCommand = "http-serve"
)

func commands() map[string]func(configFile string) {
	return map[string]func(configFile string){
		httpServeCommand: app.StartHTTPServer,
	}
}

func execute(cmd string, configFile string) {
	fmt.Println("cmd : " + cmd)
	fmt.Println("config : " + configFile)
	run, ok := commands()[cmd]
	if !ok {
		log.Fatal("invalid command")
	}

	run(configFile)
}
