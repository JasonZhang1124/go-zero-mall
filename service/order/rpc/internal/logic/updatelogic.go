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

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	//查询订单是否已存在
	var orderResult model.OrderModel
	err := l.svcCtx.DB.First(&orderResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("订单不存在")
	}

	if in.Uid != 0 {
		orderResult.Uid = uint(in.Uid)
	}
	if in.Pid != 0 {
		orderResult.Pid = uint(in.Pid)
	}
	if in.Amount != 0 {
		orderResult.Amount = uint(in.Amount)
	}
	if in.Status != 0 {
		orderResult.Status = uint(in.Status)
	}

	// 更新订单数据
	if err := l.svcCtx.DB.Model(&orderResult).Updates(orderResult).Error; err != nil {
		return nil, errors.New("订单更新失败")
	}

	return &order.UpdateResponse{}, nil
}
