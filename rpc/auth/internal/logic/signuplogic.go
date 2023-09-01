package logic

import (
	"context"

	"zero-app/rpc/auth/auth"
	"zero-app/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *auth.SignUpReq) (*auth.SignUpRes, error) {
	// todo: add your logic here and delete this line

	return &auth.SignUpRes{}, nil
}
