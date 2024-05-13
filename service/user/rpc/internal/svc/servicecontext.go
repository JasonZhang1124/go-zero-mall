package svc

import (
	"go-zero-mall/common/init_gorm"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&model.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
