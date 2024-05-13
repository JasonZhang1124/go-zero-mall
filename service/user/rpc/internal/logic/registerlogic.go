package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-mall/common/utils"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/types/user"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 判断用户是否已注册
	var userModel model.UserModel
	err := l.svcCtx.DB.Where(model.UserModel{Mobile: in.Mobile}).First(&userModel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户已存在")
	}

	newUser := model.UserModel{
		Name:     in.Name,
		Password: utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Mobile:   in.Mobile,
		Gender:   in.Gender,
	}
	result := l.svcCtx.DB.Create(&newUser)
	if result.Error != nil {
		return nil, status.Error(500, err.Error())
	}
	return &user.RegisterResponse{
		Id:     int64(newUser.ID),
		Name:   newUser.Name,
		Gender: newUser.Gender,
		Mobile: newUser.Mobile,
	}, nil
}
