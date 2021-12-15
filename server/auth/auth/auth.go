package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	OpenIdResolver OpenIdResolver
	Logger         *zap.Logger
	authpb.UnimplementedAuthServiceServer
}

type OpenIdResolver interface {
	Resolver(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	openId, err := s.OpenIdResolver.Resolver(req.Code)

	if err != nil {

		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)

	}

	return &authpb.LoginResponse{
		AccessToken: openId,
		ExpiresIn:   7200,
	}, nil
}
