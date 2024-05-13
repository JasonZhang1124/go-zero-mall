package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/rpc/types/order"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/types/pay"
	"go-zero-mall/service/user/rpc/types/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	cachePayOidPrefix = "cache:pay:oid:"
)

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}
	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{Id: in.Oid})
	if err != nil {
		return nil, err
	}
	// 查询订单是否已经创建支付
	var payResult model.PayModel
	err = l.svcCtx.DB.Where(model.PayModel{Oid: uint(in.Oid)}).First(&payResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("订单已创建")
	}
	newPay := model.PayModel{
		Uid:    uint(in.Uid),
		Oid:    uint(in.Oid),
		Amount: uint(in.Amount),
		Source: 0,
		Status: 0,
	}
	err = l.svcCtx.DB.Create(&newPay).Error
	if err != nil {
		return nil, errors.New("订单创建失败")
	}
	if newPay.ID == 0 {
		return nil, errors.New("订单创建失败")
	}

	return &pay.CreateResponse{
		Id: int64(newPay.ID),
	}, nil
}
