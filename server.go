package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hello-grpc/proto/hello"
	"log"
	"net"
)

const port = ":50051"

type SayHelloDemo struct {
	pb.UnimplementedGreeterServer
}

func (s *SayHelloDemo) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if ctx.Err() == context.Canceled {
		log.Fatal("not connectd")
		return nil, context.Canceled
	}
	fmt.Printf("HelloRequest name： %v\n", in.Name)
	return &pb.HelloReply{Message: fmt.Sprintf("hi %v\n", in.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &SayHelloDemo{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
