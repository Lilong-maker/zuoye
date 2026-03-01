package request

type OrderAdd struct {
	Name  string  `form:"name"   binding:"required"`
	Price float64 `form:"price"  binding:"required"`
	Num   int     `form:"num"  binding:"required"`
}
