package logic

import (
	"context"

	"zero-app/rpc/auth/auth"
	"zero-app/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *auth.SignInReq) (*auth.SignInRes, error) {
	generatePasswordHash(&in.Password)
	row, err := l.svcCtx.Model.FindOne(l.ctx, 3)
	if err != nil {
		return &auth.SignInRes{}, err
	}

	return &auth.SignInRes{
		Token: generateToken(row.Id),
	}, nil
}
