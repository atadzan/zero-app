package article

import (
	"context"

	"zero-app/internal/svc"
	"zero-app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogic) GetAll() (resp *types.GetArticleRes, err error) {
	// todo: add your logic here and delete this line

	return
}
