package main

import (
	"ghost-codes/waste-bin/config"
	"ghost-codes/waste-bin/server"
	"ghost-codes/waste-bin/util"
	"log"
)

func main() {
	// init config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Unable to init config,%s", err)
	}

	// init firebase
	firebaseAuth := util.SetupFirebaseClient(config.FirebaseConfigPath)

	// init server
	server, err := server.NewServer(config, firebaseAuth)
	if err != nil {
		log.Fatalf("Unable to create server,%s", err)
	}

	err = server.Start(config.ServerAddr)

	if err != nil {
		log.Fatal("error occured, Server could not start; Error:", err)
	}
}
