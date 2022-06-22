package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"xa/model"
	"xa/trans/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	UserAccountModel model.UserAccountModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	mysql := sqlx.NewMysql(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.Db,
	))

	return &ServiceContext{
		Config:           c,
		UserAccountModel: model.NewUserAccountModel(mysql, c.CacheRedis),
	}
}
