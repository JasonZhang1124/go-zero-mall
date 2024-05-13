package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/product/model"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	newProduct := model.ProductModel{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  int(in.Stock),
		Amount: uint(in.Amount),
		Status: uint(in.Status),
	}
	if err := l.svcCtx.DB.Create(&newProduct).Error; err != nil {
		return nil, errors.New("添加商品失败！")
	}
	
	return &product.CreateResponse{
		Id: int64(newProduct.ID),
	}, nil
}
