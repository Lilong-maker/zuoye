package service

import (
	"context"

	__ "zuoye/srv/dasic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedOrderServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderAdd(_ context.Context, in *__.OrderAddReq) (*__.OrderAddResp, error) {

	return &__.OrderAddResp{}, nil
}
