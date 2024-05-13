package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/pay/model"
	"gorm.io/gorm"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/types/pay"

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

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	// 查询支付是否存在
	var payResult model.PayModel
	err := l.svcCtx.DB.First(&payResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("支付不存在")
	}

	return &pay.DetailResponse{
		Id:     int64(payResult.ID),
		Uid:    int64(payResult.Uid),
		Oid:    int64(payResult.Oid),
		Amount: int64(payResult.Amount),
		Source: int64(payResult.Source),
		Status: int64(payResult.Status),
	}, nil
}
