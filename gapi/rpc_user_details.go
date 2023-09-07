package gapi

import (
	"context"
	"ghost-codes/waste-bin/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUserDetails(ctx context.Context, in *pb.CreateUserDetailsRequest) (*pb.CreateUserDetailsResponse, error) {
	uid, err := server.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	user, err := server.store.FetchUserByUID(*uid)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	res := pb.CreateUserDetailsResponse{
		User: &pb.UserDetails{
			Uid:      user.UserID,
			FullName: user.Fullname,
			Location: &pb.Location{
				Lat:  float64(user.Location.Lat),
				Long: float64(user.Location.Long),
			},
		},
	}

	return &res, nil
}

func (server *Server) FetchUserDetails(ctx context.Context, in *pb.FetchUserParams) (*pb.FetchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserDetails not implemented")
}
