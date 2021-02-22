package logic

import (
	"context"

	"bookstore/rpc/check/internal/svc"
	check "bookstore/rpc/check/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {

	book, err := l.svcCtx.Model.FindOne(in.Book)
	if err != nil {
		return nil, err
	}
	return &check.CheckResp{
		Price: book.Price,
		Found: true,
	}, nil
}
