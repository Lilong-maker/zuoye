package service

import (
	"context"
	"zuoye/srv/dasic/config"
	"zuoye/srv/handler/model"

	__ "zuoye/srv/dasic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedOrderServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderAdd(_ context.Context, in *__.OrderAddReq) (*__.OrderAddResp, error) {

	var order model.Order
	err := order.FindOrder(config.DB, in.Name)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "参数错误",
			Code: 400,
		}, nil
	}
	m := model.Order{
		Name:  in.Name,
		Price: float64(in.Price),
		Num:   int(in.Num),
	}
	err = m.OrderAdd(config.DB)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "订单添加失败",
			Code: 400,
		}, nil
	}
	return &__.OrderAddResp{
		Msg:  "订单添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderDelete(_ context.Context, in *__.OrderDeleteReq) (*__.OrderDeleteResp, error) {

	return &__.OrderDeleteResp{
		Msg:  "订单删除成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderUpdate(_ context.Context, in *__.OrderUpdateReq) (*__.OrderUpdateResp, error) {

	return &__.OrderUpdateResp{
		Msg:  "订单修改成功",
		Code: 200,
	}, nil
}
