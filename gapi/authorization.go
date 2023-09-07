package gapi

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationType   = "Bearer"
)

func (server *Server) authorization(ctx context.Context) (*string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)

	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := fields[0]
	if authType != authorizationType {
		return nil, fmt.Errorf("invalid authorization type")
	}

	token, err := server.firebaseAuth.VerifyIDToken(ctx, fields[1])

	if err != nil {

		return nil, err
	}

	if err != nil {

		return nil, err
	}

	return &token.UID, nil
}
