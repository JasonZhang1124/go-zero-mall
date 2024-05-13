package model

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	Name   string `gorm:"comment:产品名称"`
	Desc   string `gorm:"comment:产品描述"`
	Stock  int    `gorm:"comment:产品库存"`
	Amount uint   `gorm:"comment:产品金额"`
	Status uint   `gorm:"comment:产品状态"`
}
