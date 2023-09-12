package gapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rpip/paystack-go"
)

type TransactionReqResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

func (server *Server) initializeTransaction(email string, amountCents int) (*string, error) {
	client := paystack.NewClient(server.config.PaystackSK, http.DefaultClient)

	response, err := client.Transaction.Initialize(&paystack.TransactionRequest{
		Currency:    "GHS",
		Email:       email,
		Amount:      float32(amountCents),
		CallbackURL: "www.google.com",
	})

	fmt.Println(err)
	fmt.Println(response)
	if err != nil {
		return nil, err
	}
	res := TransactionReqResponse{}

	b, err := json.Marshal(response)
	if err != nil {
		return nil, errors.New("Paystack response error")
	}
	err = json.Unmarshal(b, &res)

	if err != nil {
		return nil, errors.New("Paystack response error")
	}

	return &res.AuthorizationURL, nil
}

func (server *Server) paymentCallback() error {
	return nil
}
