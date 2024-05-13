package logic

import (
	"context"
	"errors"
	"go-zero-mall/common/utils"
	"go-zero-mall/service/user/model"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 判断用户是否存在
	var result model.UserModel
	err := l.svcCtx.DB.Where(model.UserModel{Mobile: in.Mobile}).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(500, "用户不存在")
	}
	// 判断密码
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != result.Password {
		return nil, status.Error(500, "密码错误")
	}
	return &user.LoginResponse{
		Id:     int64(result.ID),
		Name:   result.Name,
		Gender: result.Gender,
		Mobile: result.Mobile,
	}, nil
}
