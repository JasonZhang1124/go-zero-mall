package model

import "gorm.io/gorm"

type OrderModel struct {
	gorm.Model
	Uid    uint `gorm:"comment:用户ID"`
	Pid    uint `gorm:"comment:产品ID"`
	Amount uint `gorm:"comment:订单价格"`
	Status uint `gorm:"comment:订单状态"`
}
