package dao

import (
	"context"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/conf"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/model"
	"github.com/MenciusCheng/superman/util/dragons/proxy"
)

// Dao represents data access object
type Dao struct {
	c  *conf.Config
	db *proxy.SQL
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c:  c,
		db: proxy.InitSQL("bookmark_manager"),
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

func (d *Dao) GetLinkList(ctx context.Context) (infos []model.Link, err error) {
	err = d.db.Master(ctx).Find(&infos).Error
	return
}

func (d *Dao) CreateLinkList(ctx context.Context, infos []model.Link) (err error) {
	err = d.db.Master(ctx).CreateInBatches(&infos, 500).Error
	return
}
