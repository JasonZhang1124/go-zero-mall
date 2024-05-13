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

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
	//查询订单是否已存在
	var orderResult model.OrderModel
	err := l.svcCtx.DB.First(&orderResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("订单不存在")
	}

	return &order.DetailResponse{
		Id:     int64(orderResult.ID),
		Uid:    int64(orderResult.Uid),
		Pid:    int64(orderResult.Pid),
		Amount: int64(orderResult.Amount),
		Status: int64(orderResult.Status),
	}, nil
}
