package logic

import (
	"context"
	"go-zero-mall/service/order/rpc/types/order"

	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	result, err := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     result.Id,
		Uid:    result.Uid,
		Pid:    result.Pid,
		Amount: result.Amount,
		Status: result.Status,
	}, nil
}
