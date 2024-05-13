package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/common/init_gorm"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/orderclient"
	"go-zero-mall/service/pay/rpc/internal/config"
	"go-zero-mall/service/user/rpc/userclient"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	_ = db.AutoMigrate(&model.OrderModel{})
	return &ServiceContext{
		Config:   c,
		DB:       db,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
