package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/product/model"
	"gorm.io/gorm"

	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/types/product"

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

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	//查询商品是否存在
	var result model.ProductModel
	err := l.svcCtx.DB.First(&result, in.Id).Error
	// 不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("商品不存在")
	}

	// 删除
	if err = l.svcCtx.DB.Delete(&result, in.Id).Error; err != nil {
		return nil, errors.New("商品删除失败")
	}
	return &product.RemoveResponse{}, nil
}
