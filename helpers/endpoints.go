package helpers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type EncryptService interface {
	Encrypt( context.Context, string, string ) (string, error)
	Decrypt( context.Context, string, string ) (string, error)
}

// MakeEncryptEndpoint forms endpoint for request/response of encrypt function
func MakeEncryptEndpoint(svc EncryptService) endpoint.Endpoint {
	fmt.Println("MakeEncryptEndpoint called")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("MakeEncryptEndpoint Starts")
		req := request.(EncryptRequest)
		message, err := svc.Encrypt(ctx, req.Key, req.Text)
		if err != nil {
			return EncryptResponse{message, err.Error()}, nil
		}
		fmt.Println("MakeEncryptEndpoint Ends")
		return EncryptResponse{message, ""}, nil
	}
}
// MakeDecryptEndpoint forms endpoint for request/response of decrypt function
func MakeDecryptEndpoint(svc EncryptService) endpoint.Endpoint {
	fmt.Println("MakeDecryptEndpoint called")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DecryptRequest)
		text, err := svc.Decrypt(ctx, req.Key, req.Message)
		if err != nil {
			return DecryptResponse{text, err.Error()}, nil
		}
		return DecryptResponse{text, ""}, nil
	}
}
