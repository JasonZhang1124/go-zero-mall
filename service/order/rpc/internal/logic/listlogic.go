package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/user/rpc/types/user"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	//按用户id查询订单
	var orderList []model.OrderModel
	err = l.svcCtx.DB.Where(model.OrderModel{Uid: uint(in.Uid)}).Find(&orderList).Error
	if err != nil {
		return nil, errors.New("查询发送错误")
	}
	orderDetailList := make([]*order.DetailResponse, 0)
	for _, item := range orderList {
		orderDetailList = append(orderDetailList, &order.DetailResponse{
			Id:     int64(item.ID),
			Uid:    int64(item.Uid),
			Pid:    int64(item.Pid),
			Amount: int64(item.Amount),
			Status: int64(item.Status),
		})
	}
	return &order.ListResponse{
		Data: orderDetailList,
	}, nil
}
