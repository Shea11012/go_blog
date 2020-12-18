package service

import (
	"context"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}

	return svc
}
