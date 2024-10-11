package main

import (
	"boiler-plate/init/cmd"
	"flag"
)

var configPathFlag = flag.String("config", "./init/config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
