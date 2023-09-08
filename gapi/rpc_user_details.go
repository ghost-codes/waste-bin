package gapi

import (
	"context"
	db "ghost-codes/waste-bin/db/models"
	"ghost-codes/waste-bin/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUserDetails(ctx context.Context, in *pb.CreateUserDetailsRequest) (*pb.CreateUserDetailsResponse, error) {
	uid, err := server.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	user := db.UserDetails{
		UserID:   *uid,
		Fullname: in.FullName,
		Location: db.Location{
			Lat:  float32(in.Location.Lat),
			Long: float32(in.Location.Long),
			Name: in.Location.Name,
		},
	}
	err = server.store.CreateUser(&user)

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
				Name: user.Location.Name,
			},
		},
	}

	return &res, nil
}

func (server *Server) FetchUserDetails(ctx context.Context, in *pb.FetchUserParams) (*pb.FetchUserResponse, error) {
	uid, err := server.authorization(ctx)

	user, err := server.store.FetchUserByUID(*uid)
	if err != nil {
		return &pb.FetchUserResponse{
			IsSignUpd: false,
		}, nil
	}

	res := pb.FetchUserResponse{
		User: &pb.UserDetails{
			Uid:      user.UserID,
			FullName: user.Fullname,
			Location: &pb.Location{
				Lat:  float64(user.Location.Lat),
				Long: float64(user.Location.Long),
				Name: user.Location.Name,
			},
		},
		IsSignUpd: true,
	}
	return &res, nil
}
