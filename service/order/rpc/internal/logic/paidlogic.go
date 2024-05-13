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

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *order.PaidRequest) (*order.PaidResponse, error) {
	//查询订单是否已存在
	var orderResult model.OrderModel
	err := l.svcCtx.DB.First(&orderResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("订单不存在")
	}

	orderResult.Status = 1
	err = l.svcCtx.DB.Model(&orderResult).Select("status").Updates(orderResult).Error
	if err != nil {
		return nil, errors.New("更新订单状态失败")
	}

	return &order.PaidResponse{}, nil
}
