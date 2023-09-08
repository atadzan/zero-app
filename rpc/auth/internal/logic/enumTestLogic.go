package logic

import (
	"context"

	"zero-app/rpc/auth/auth"
	"zero-app/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnumTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnumTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnumTestLogic {
	return &EnumTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnumTestLogic) EnumTest(in *auth.EnumReq) (*auth.EnumRes, error) {
	// todo: add your logic here and delete this line

	return &auth.EnumRes{}, nil
}
