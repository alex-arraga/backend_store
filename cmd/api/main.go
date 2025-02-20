package main

import "log"

func main() {
	config, err := LoadAppConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	StartServer(config)
}
