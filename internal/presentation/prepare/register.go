package prepare

import (
	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
)

type CommandServer struct {
	Server *grpc.Server
}

func NewCommandServer(category pb.CategoryCommandServer, product pb.ProductCommandServer) *CommandServer {
	server := grpc.NewServer()
	pb.RegisterCategoryCommandServer(server, category)
	pb.RegisterProductCommandServer(server, product)
	return &CommandServer{
		Server: server,
	}
}
