package service

import (
	"context"

	pb "lucy/api/engine/service/v1"
)

type EngineService struct {
	pb.UnimplementedEngineServer
}

func NewEngineService() *EngineService {
	return &EngineService{}
}

func (s *EngineService) CreateOrder(ctx context.Context, req *pb.AddOrderReq) (*pb.AddOrderReply, error) {
	return &pb.AddOrderReply{}, nil
}
func (s *EngineService) CancelOrder(ctx context.Context, req *pb.CancelOrderReq) (*pb.CancelOrderReply, error) {
	return &pb.CancelOrderReply{}, nil
}
func (s *EngineService) AddSymbol(ctx context.Context, req *pb.AddSymbolReq) (*pb.AddSymbolReply, error) {
	return &pb.AddSymbolReply{}, nil
}
func (s *EngineService) GetEngine(ctx context.Context, req *pb.CloseSymbolReq) (*pb.CloseSymbolReply, error) {
	return &pb.CloseSymbolReply{}, nil
}
