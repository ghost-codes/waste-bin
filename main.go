package main

import (
	"context"
	"ghost-codes/waste-bin/config"
	"ghost-codes/waste-bin/gapi"
	"ghost-codes/waste-bin/pb"
	"ghost-codes/waste-bin/util"
	"log"
	"net"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Unable to init config,%s", err)
	}

	// init firebase
	firebaseAuth := util.SetupFirebaseClient(config.FirebaseConfigPath)

	go runGRPCGateway(config, firebaseAuth)
	runGRPCServer(config, firebaseAuth)

}

func runGRPCServer(config config.Config, authClient *auth.Client) {
	// init server
	server, err := gapi.NewServer(config, authClient)
	if err != nil {
		log.Fatalf("Unable to create server,%s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWasteBinServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("Unable to create lister,%s", err)
	}

	log.Printf("gRPC server started at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server,%s", err)
	}
}
func runGRPCGateway(config config.Config, authClient *auth.Client) {
	// init server
	server, err := gapi.NewServer(config, authClient)
	if err != nil {
		log.Fatalf("Unable to create server,%s", err)
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())

	err = pb.RegisterWasteBinHandlerServer(ctx, grpcMux, server)
	defer cancel()

	if err != nil {
		log.Fatalf("cannot registion handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HttpServerAddress)
	if err != nil {
		log.Fatalf("Unable to create lister,%s", err)
	}

	log.Printf("gRPC server started at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP gateway server,%s", err)
	}
}
