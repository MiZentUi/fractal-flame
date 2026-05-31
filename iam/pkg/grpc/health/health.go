package health

import (
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func Register(server *grpc.Server) {
	health := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, health)

	health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	health.SetServingStatus(iamv1.IAMService_ServiceDesc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
}
