package health

import (
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func Register(server *grpc.Server) {
	health := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, health)

	health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
}
