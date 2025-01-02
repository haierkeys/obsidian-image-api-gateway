package service

import (
	"context"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	svc.dao = dao.New(global.DBEngine)
	return svc
}

func (svc *Service) Ctx() context.Context {
	return svc.ctx
}
