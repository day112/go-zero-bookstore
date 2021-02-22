package logic

import (
	"bookstore/rpc/add/internal/svc"
	add "bookstore/rpc/add/pb"
	"bookstore/rpc/check/checker"
	"bookstore/rpc/model/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// 4. 在 AddLogic 添加逻辑中 调用 check 服务
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: in.Book,
	})
	if err == model.ErrNotFound {
		_, err = l.svcCtx.Model.Insert(model.Book{
			Book:  in.Book,
			Price: in.Price,
		})
		if err != nil {
			return nil, err
		}
		return &add.AddResp{
			Ok: true,
		}, nil
	}else if err != nil {
		return &add.AddResp{
			Ok: false,
		}, nil
	}
	if resp.Found {
		return &add.AddResp{
			Ok: false,
		}, nil
	}
	return &add.AddResp{
		Ok: false,
	}, nil
}
