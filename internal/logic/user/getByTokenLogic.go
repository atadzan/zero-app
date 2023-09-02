package user

import (
	"context"

	"zero-app/internal/svc"
	"zero-app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByTokenLogic {
	return &GetByTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByTokenLogic) GetByToken() (resp *types.GetProfileRes, err error) {

	return
}
