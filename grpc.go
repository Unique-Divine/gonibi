package gonibi

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGRPCConnection(
	grpcUrl string, grpcInsecure bool, timeoutSeconds int64,
) (*grpc.ClientConn, error) {
	var creds credentials.TransportCredentials
	if grpcInsecure {
		creds = insecure.NewCredentials()
	} else {
		creds = credentials.NewTLS(&tls.Config{})
	}

	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
	}
	timeout := time.Duration(timeoutSeconds) * time.Second
	ctx, cancel := context.WithTimeout(
		context.Background(), timeout,
	)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcUrl, options...)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: Cannot connect to gRPC endpoint %s\n", err, grpcUrl)
	}

	return conn, nil
}

// var _ NibiruClient = (*GrpcClient)(nil)

// func NewGrpcClient(grpcConn *grpc.ClientConn) NibiruClient {
// }