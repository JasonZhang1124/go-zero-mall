package logic

import (
	"context"
	"errors"
	"go-zero-mall/service/user/model"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 查询用户是否已存在
	var result model.UserModel
	err := l.svcCtx.DB.First(&result, in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(500, "用户不存在")
	}
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &user.UserInfoResponse{
		Id:     int64(result.ID),
		Name:   result.Name,
		Gender: result.Gender,
		Mobile: result.Mobile,
	}, nil
}
