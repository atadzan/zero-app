package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-app/internal/svc"
)

type DocsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocsLogic {
	return &DocsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocsLogic) Docs() error {
	// todo: add your logic here and delete this line

	return nil
}
