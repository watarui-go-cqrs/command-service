package prepare

import (
	"command-service/internal/infrastructure/sqlboiler/handler"
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func CommandServiceLifecycle(lifecycle fx.Lifecycle, server *CommandServer) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := handler.DBConnect(); err != nil {
					panic(err)
				}
				port := 8082
				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
				if err != nil {
					return err
				}
				reflection.Register(server.Server)
				go func() {
					log.Printf("gRPC server is listening on port %d", port)
					server.Server.Serve(listener)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				server.Server.GracefulStop()
				log.Println("gRPC server stopped gracefully")
				return nil
			},
		},
	)
}
