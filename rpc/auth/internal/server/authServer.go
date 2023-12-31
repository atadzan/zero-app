// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package server

import (
	"context"

	"zero-app/rpc/auth/auth"
	"zero-app/rpc/auth/internal/logic"
	"zero-app/rpc/auth/internal/svc"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedAuthServer
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServer) SignIn(ctx context.Context, in *auth.SignInReq) (*auth.SignInRes, error) {
	l := logic.NewSignInLogic(ctx, s.svcCtx)
	return l.SignIn(in)
}

func (s *AuthServer) SignUp(ctx context.Context, in *auth.SignUpReq) (*auth.SignUpRes, error) {
	l := logic.NewSignUpLogic(ctx, s.svcCtx)
	return l.SignUp(in)
}

func (s *AuthServer) EnumTest(ctx context.Context, in *auth.EnumReq) (*auth.EnumRes, error) {
	l := logic.NewEnumTestLogic(ctx, s.svcCtx)
	return l.EnumTest(in)
}
