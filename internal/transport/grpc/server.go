package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	urlgrpc "github.com/Mimist-Illusionard/tny/api/v1/grpc"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const shutdownTimeout = 5 * time.Second

func Serve(ctx context.Context, port string, service Shortener) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("listen gRPC on port %s: %w", port, err)
	}

	server := googlegrpc.NewServer()
	urlgrpc.RegisterURLShortenerServer(server, NewHandler(service))
	reflection.Register(server)

	errCh := make(chan error, 1)
	go func() {
		log.Printf("grpc server listening on :%s", port)
		errCh <- server.Serve(listener)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("serve gRPC: %w", err)
		}
		return nil
	case <-ctx.Done():
		stopDone := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopDone)
		}()

		select {
		case <-stopDone:
		case <-time.After(shutdownTimeout):
			server.Stop()
		}

		if err := <-errCh; err != nil {
			return fmt.Errorf("stop gRPC server: %w", err)
		}
		return nil
	}
}
