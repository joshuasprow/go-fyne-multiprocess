package main

import (
	"context"
	"log"
	"net"

	pb "github.com/joshuasprow/go-fyne-multiprocess/api"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer

	exit chan struct{}
}

func newServer() *server {
	return &server{exit: make(chan struct{})}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("hello from: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	log.Printf("goodbye from: %v", in.GetName())

	s.exit <- struct{}{}

	return &pb.GoodbyeReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	log.SetPrefix("server: ")

	srv := newServer()
	var s *grpc.Server

	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		s = grpc.NewServer()

		log.Printf("Server listening on %s", port)

		pb.RegisterGreeterServer(s, srv)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-srv.exit

	s.GracefulStop()
}
