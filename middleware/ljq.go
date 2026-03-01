package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func UnaryServerLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// 获取客户端地址
	var clientAddr string
	if p, ok := peer.FromContext(ctx); ok {
		clientAddr = p.Addr.String()
	}
	// 调⽤实际 handler
	resp, err := handler(ctx, req)
	// 记录⽇志：⽅法、客户端、耗时、错误（如果有）
	log.Printf("[Server Unary] method=%s client=%s duration=%s error=%v",
		info.FullMethod, clientAddr, time.Since(start), err)
	return resp, err
}
