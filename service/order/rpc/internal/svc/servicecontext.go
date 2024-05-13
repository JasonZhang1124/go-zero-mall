package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/common/init_gorm"
	"go-zero-mall/service/order/rpc/internal/config"
	"go-zero-mall/service/product/rpc/productclient"
	"go-zero-mall/service/user/rpc/userclient"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		DB:         db,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
