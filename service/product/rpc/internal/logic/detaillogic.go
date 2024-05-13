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

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	//查询商品是否存在
	var result model.ProductModel
	err := l.svcCtx.DB.First(&result, in.Id).Error
	// 不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("商品不存在")
	}
	return &product.DetailResponse{
		Id:     int64(result.ID),
		Name:   result.Name,
		Desc:   result.Desc,
		Stock:  int64(result.Stock),
		Amount: int64(result.Amount),
		Status: int64(result.Status),
	}, nil
}
