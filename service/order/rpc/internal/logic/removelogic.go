package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/model"
	"gorm.io/gorm"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	//查询订单是否已存在
	var orderResult model.OrderModel
	err := l.svcCtx.DB.First(&orderResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("订单不存在")
	}

	if err = l.svcCtx.DB.Delete(&orderResult).Error; err != nil {
		return nil, errors.New("订单删除失败")
	}
	return &order.RemoveResponse{}, nil
}
