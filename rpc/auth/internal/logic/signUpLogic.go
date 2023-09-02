package logic

import (
	"context"
	"zero-app/rpc/auth/model"

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
	generatePasswordHash(&in.Password)
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.TestUsers{
		Lastname: in.Lastname,
		Username: in.Username,
		Password: in.Password,
		Age:      in.Age,
	})
	if err != nil {
		return &auth.SignUpRes{}, err
	}

	return &auth.SignUpRes{
		Message: "success",
	}, nil
}
