package gapi

import (
	"fmt"
	"ghost-codes/waste-bin/config"
	"ghost-codes/waste-bin/db"
	models "ghost-codes/waste-bin/db/models"
	"ghost-codes/waste-bin/pb"

	"firebase.google.com/go/auth"
)

type Server struct {
	pb.UnimplementedWasteBinServer
	config       config.Config
	store        *models.Store
	firebaseAuth *auth.Client
}

func NewServer(config config.Config, firebaseAuth *auth.Client) (*Server, error) {
	dbIns, err := db.NewGorm(config.DBSource())
	if err != nil {
		return nil, fmt.Errorf("unable to initialize db,%s", err)
	}
	server := Server{
		config:       config,
		firebaseAuth: firebaseAuth,
		store: &models.Store{
			DB: dbIns,
		},
	}

	return &server, nil
}
