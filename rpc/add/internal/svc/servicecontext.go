package svc

import (
	"bookstore/rpc/add/internal/config"
	"bookstore/rpc/check/checker"
	"bookstore/rpc/model/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	c     config.Config
	Model *model.BookModel
	// 3. ServiceContext 添加服务上下文配置
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
