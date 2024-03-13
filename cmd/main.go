package main

import (
	"flag"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "this is config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
	flag.Parse()

	execute(flag.Args()[0], configFile)
}
