package main

import (
	"bytes"
	"context"
	"fmt"
	"ghost-codes/waste-bin/config"
	"ghost-codes/waste-bin/gapi"
	"ghost-codes/waste-bin/pb"
	"ghost-codes/waste-bin/util"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

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
	if config.AppEnv == "development" {
		getDummyIdToken(config, firebaseAuth)
	}

	go runGRPCGateway(config, firebaseAuth)
	runGRPCServer(config, firebaseAuth)

}

func getDummyIdToken(config config.Config, authClient *auth.Client) {
	token, err := authClient.CustomToken(context.Background(), config.DevUID)
	if err != nil {
		fmt.Println(err)
	}
	jsonBody := []byte(fmt.Sprintf(`{"token":"%s","returnSecureToken":%v}`, token, true))
	bodyReader := bytes.NewReader(jsonBody)

	requestUrl := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", config.ProjectAPIKey)
	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

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

	log.Printf("Http server started at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP gateway server,%s", err)
	}
}
