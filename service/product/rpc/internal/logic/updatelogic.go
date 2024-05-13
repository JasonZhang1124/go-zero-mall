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

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	//查询商品是否存在
	var result model.ProductModel
	err := l.svcCtx.DB.First(&result, in.Id).Error
	// 不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("商品不存在")
	}
	if in.Name != "" {
		result.Name = in.Name
	}
	if in.Desc != "" {
		result.Desc = in.Desc
	}
	if in.Stock != 0 {
		result.Stock = int(in.Stock)
	}
	if in.Amount != 0 {
		result.Amount = uint(in.Amount)
	}
	if in.Status != 0 {
		result.Status = uint(in.Status)
	}
	if err := l.svcCtx.DB.Model(&result).Updates(result).Error; err != nil {
		return nil, errors.New("更新商品失败")
	}
	return &product.UpdateResponse{}, nil
}
