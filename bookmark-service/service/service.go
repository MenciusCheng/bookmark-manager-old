package service

import (
	"context"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/conf"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/dao"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/manager"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/model"
)

type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

func New(c *conf.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
	if s.mgr != nil {
		s.mgr.Close()
	}
}

func (s *Service) GetLinkList(ctx context.Context) (infos []model.Link, err error) {
	return s.dao.GetLinkList(ctx)
}
