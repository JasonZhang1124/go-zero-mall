package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/rpc/types/order"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/user/rpc/types/user"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}
	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询支付是否存在
	var payResult model.PayModel
	err = l.svcCtx.DB.First(&payResult, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(100, "支付不存在")
	}
	// 支付金额与订单金额不符
	if uint(in.Amount) != payResult.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	payResult.Source = uint(in.Source)
	payResult.Status = uint(in.Status)

	err = l.svcCtx.DB.Save(payResult).Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil
}
