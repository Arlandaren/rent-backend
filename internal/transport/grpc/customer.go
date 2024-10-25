package grpc

import (
	desc "customer_service/pkg/grpc/customer_v1"
)

type customerServer struct {
	desc.UnimplementedCustomerServiceServer
}
