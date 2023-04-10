package service

import (
	"context"
	"github.com/shopspring/decimal"
	"lucy/app/engine/internal/biz"

	pb "lucy/api/engine/service/v1"
)

type EngineService struct {
	pb.UnimplementedEngineServer
	engine *biz.Engine
}

func NewEngineService(engine *biz.Engine) *EngineService {
	return &EngineService{
		engine: engine,
	}
}

func (s *EngineService) CreateOrder(ctx context.Context, req *pb.AddOrderReq) (*pb.AddOrderReply, error) {
	price, err := decimal.NewFromString(req.Order.Price)
	if err != nil {
		return &pb.AddOrderReply{Reply: &pb.BasicReply{Code: 1, Message: ""}}, err
	}
	quantity, err := decimal.NewFromString(req.Order.Quantity)
	if err != nil {
		return &pb.AddOrderReply{Reply: &pb.BasicReply{Code: 1, Message: ""}}, err
	}
	order := &biz.Order{
		OrderId:  req.Order.OrderId,
		UserId:   req.Order.UserId,
		Symbol:   req.Order.Symbol,
		Price:    price,
		Quantity: quantity,
		Side:     req.Order.Side,
		Type:     req.Order.Type,
	}
	if err := s.engine.Add(order); err != nil {
		return &pb.AddOrderReply{Reply: &pb.BasicReply{Code: 1, Message: ""}}, err
	}
	return &pb.AddOrderReply{}, nil
}
func (s *EngineService) CancelOrder(ctx context.Context, req *pb.CancelOrderReq) (*pb.CancelOrderReply, error) {
	// 从订单簿中删除
	if err := s.engine.Cancel(req.Symbol, req.OrderId); err != nil {
		return nil, err
	}
	return &pb.CancelOrderReply{}, nil
}
func (s *EngineService) AddSymbol(ctx context.Context, req *pb.AddSymbolReq) (*pb.AddSymbolReply, error) {
	if err := s.engine.Open(req.Symbol); err != nil {
		return nil, err
	}
	return &pb.AddSymbolReply{}, nil
}
func (s *EngineService) GetEngine(ctx context.Context, req *pb.CloseSymbolReq) (*pb.CloseSymbolReply, error) {
	if err := s.engine.Close(req.Symbol); err != nil {
		return nil, err
	}
	return &pb.CloseSymbolReply{}, nil
}
