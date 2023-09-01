package article

import (
	"context"

	"zero-app/internal/svc"
	"zero-app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllArticlesLogic {
	return &GetAllArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllArticlesLogic) GetAllArticles() (resp []types.GetArticleRes, err error) {
	// todo: add your logic here and delete this line

	return
}
