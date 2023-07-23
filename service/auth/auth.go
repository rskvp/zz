package auth

import (
	"context"
	pb "zz/protos/gen/auth"

	"zz/service/auth/jwt"
	"zz/service/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authAPIServer struct {
	pb.UnimplementedAuthAPIServer
	jwt *jwt.Service
}

func NewAPI(jwt *jwt.Service) *authAPIServer {
	if jwt == nil {
		logger.I.Panic("jwt is nil")
	}
	return &authAPIServer{
		jwt: jwt,
	}
}

func (s *authAPIServer) GetJWT(ctx context.Context, req *pb.GetJWTRequest) (*pb.GetJWTResponse, error) {
	logger.I.Debug("GetJWT called", zap.Any("req", req))
	userID, err := Validate(ctx, req.GetAccount())
	if err != nil {
		logger.I.Error("user failed to authenticate", zap.Error(err), zap.Any("req", req))
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated: %s", err)
	}
	token := s.jwt.CreateToken(userID)
	logger.I.Debug("GetJWT success", zap.Any("token", token))
	return &pb.GetJWTResponse{
		Token: s.jwt.CreateToken(userID),
	}, nil
}
