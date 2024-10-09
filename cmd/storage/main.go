package main

import (
	"flag"

	"github.com/Utro-tvar/Storage/internal/config"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")

	flag.Parse()

	cfg := config.MustLoad(*configPath)

	_ = cfg
}
