package gapi

import (
	"context"
	"fmt"
	db "ghost-codes/waste-bin/db/models"
	"ghost-codes/waste-bin/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewRequest(request db.Request) *pb.Request {
	return &pb.Request{
		Id:            int32(request.ID),
		UserId:        request.UserID,
		Payment:       NewPayment(request.Payment),
		State:         pb.REQUEST_STATE(pb.REQUEST_STATE_value[string(request.State)]),
		DateOfDropoff: timestamppb.New(request.DateOfDropOff),
		Bin:           NewBin(request.Bin),
		Type:          pb.RQUEST_TYPE(pb.RQUEST_TYPE_value[string(request.Type)]),
	}
}

func NewPayment(payment *db.Payment) *pb.Payment {
	if payment == nil {
		return nil
	}

	return &pb.Payment{
		Id:     int32(payment.ID),
		Amount: int32(payment.Amount),
		Date:   timestamppb.New(payment.Date),
	}
}

func NewBin(bin *db.Bin) *pb.Bin {
	if bin == nil {
		return nil
	}
	return &pb.Bin{
		Id:     uint32(bin.ID),
		Status: pb.BinStatus(pb.BinStatus_value[string(bin.Status)]),
		Type:   pb.BinType(pb.BinType_value[string(bin.Type)]),
	}
}

func (server *Server) MakeRequest(ctx context.Context, payload *pb.CreateRequestPayload) (*pb.CreateRequestResponse, error) {
	uid, err := server.authorization(ctx)
	if err != nil {
		return nil, err
	}
	var bin *db.Bin = nil

	if payload.BinType != nil {
		bin = &db.Bin{
			Type:   db.BinType(payload.BinType.String()),
			UserID: *uid,
		}
	}

	var special_pickup *db.SpecialPickup = nil

	if payload.Type.String() == string(db.SPECIALPICKUP) {
		special_pickup = &db.SpecialPickup{
			UserID:       *uid,
			DateOfPickup: payload.Date.AsTime(),
		}
	}

	fmt.Println(pb.RQUEST_TYPE_name[int32(payload.Type)])
	fmt.Println(payload.Type)

	request := db.Request{
		UserID:        *uid,
		Type:          db.RequestType(pb.RQUEST_TYPE_name[int32(payload.Type)]),
		Bin:           bin,
		SpecialPickup: special_pickup,
		// DateOfDropOff: ,
	}

	err = server.store.CreateRequest(&request)
	if err != nil {
		return nil, err
	}

	user, err := server.firebaseAuth.GetUser(ctx, *uid)
	if err != nil {
		return nil, err
	}
	paymentUrl, err := server.initializeTransaction(user.Email, 1000)
	return &pb.CreateRequestResponse{
		Data:       NewRequest(request),
		PaymentUrl: *paymentUrl,
	}, nil
}
