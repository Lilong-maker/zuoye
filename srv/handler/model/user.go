package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name  string  `grom:"type:varchar(30)"`
	Price float64 `grom:"type:decimal(10,2)"`
	Num   int     `grom:"type:int"`
}
type Orders struct {
	gorm.Model
	Name  string  `grom:"type:varchar(30)"`
	Price float64 `grom:"type:decimal(10,2)"`
	Num   int     `grom:"type:int"`
}
