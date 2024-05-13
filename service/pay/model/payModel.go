package model

import "gorm.io/gorm"

type PayModel struct {
	gorm.Model
	Uid    uint `gorm:"primaryKey;default:0;not null;comment:用户ID"`
	Oid    uint `gorm:"primaryKey;default:0;not null;comment:订单ID"`
	Amount uint `gorm:"default:0;not null;comment:产品金额"`
	Source uint `gorm:"default:0;not null;comment:支付方式"`
	Status uint `gorm:"default:0;not null;comment:支付状态"`
}
