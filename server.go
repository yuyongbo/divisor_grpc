package main

import (
	"log"
	"net"
	// Change this for your own project
	"fmt"
	
	context "golang.org/x/net/context"
	"divisor_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
)

//   在 主要 功能,注册一个服务器类型将处理请求。 然后开始gRPC服务器。

type server struct{}

func main() {


	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//   声明 计算 处理函数。 这使得 服务器 符合自动生成的类型 pb.GCDServiceServer 接口。

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a,b)
	return &pb.GCDResponse{Result: a}, nil
}
