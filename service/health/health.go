package health

import (
	"context"
	"time"

	pb "zz/protos/gen/health"

	"zz/service/logger"

	"go.uber.org/zap"
)

type health struct {
	pb.UnimplementedHealthServer
}

func New() *health {
	return &health{}
}

func (h *health) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (h *health) Watch(
	req *pb.HealthCheckRequest,
	stream pb.Health_WatchServer,
) error {
	ctx := stream.Context()
	ticker := time.NewTicker(10 * time.Second)

	go func(ctx context.Context) {
		for {
			if err := stream.Send(&pb.HealthCheckResponse{
				Status: pb.HealthCheckResponse_SERVING,
			}); err != nil {
				logger.I.Error("healthcheck send failed", zap.Error(err))
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}(ctx)

	return nil
}
