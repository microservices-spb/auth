package rpc

import (
	"context"
	"fmt"
	"github.com/microservices-spb/auth/pkg/auth"
)

type Handler struct {
	auth.UnimplementedAuthServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	logpass := map[string]string{
		"qwerty": "qwerty",
		"login":  "login",
	}
	pass, ok := logpass[in.Username]
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}
	if pass != in.Password {
		return nil, fmt.Errorf("invalid password")
	}
	return &auth.LoginOut{
		Token: "success",
	}, nil
}
