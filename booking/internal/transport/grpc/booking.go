package grpc

import (
	"context"
	"service/internal/service"

	desc "service/pkg/grpc/booking_v1"
)

type Server struct {
	desc.BookingServiceServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}

func (s *Server) NewBooking(ctx context.Context, req *desc.NewBookingRequest) (*desc.Booking, error) {
	return s.Service.New(ctx, req)
}

func (s *Server) BeginBooking(ctx context.Context, req *desc.BeginBookingRequest) (*desc.BeginBookingResponse, error) {
	return s.Service.Begin(ctx, req)
}

func (s *Server) FinishBooking(ctx context.Context, req *desc.FinishBookingRequest) (*desc.FinishBookingResponse, error) {
	return s.Service.Finish(ctx, req)
}

func (s *Server) CancelBooking(ctx context.Context, req *desc.CancelBookingRequest) (*desc.CancelBookingResponse, error) {
	return s.Service.Cancel(ctx, req)
}

func (s *Server) UpdateBooking(ctx context.Context, req *desc.UpdateBookingRequest) (*desc.UpdateBookingResponse, error) {
	return s.Service.Update(ctx, req)
}
