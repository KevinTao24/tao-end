package service

import (
	"context"

	pb "tao-end/api/fund/v1"
)

type FundService struct {
	pb.UnimplementedFundServer
}

func NewFundService() *FundService {
	return &FundService{}
}

func (s *FundService) CreateFund(ctx context.Context, req *pb.CreateFundRequest) (*pb.CreateFundReply, error) {
	return &pb.CreateFundReply{}, nil
}
func (s *FundService) UpdateFund(ctx context.Context, req *pb.UpdateFundRequest) (*pb.UpdateFundReply, error) {
	return &pb.UpdateFundReply{}, nil
}
func (s *FundService) DeleteFund(ctx context.Context, req *pb.DeleteFundRequest) (*pb.DeleteFundReply, error) {
	return &pb.DeleteFundReply{}, nil
}
func (s *FundService) GetFund(ctx context.Context, req *pb.GetFundRequest) (*pb.GetFundReply, error) {
	return &pb.GetFundReply{}, nil
}
func (s *FundService) ListFund(ctx context.Context, req *pb.ListFundRequest) (*pb.ListFundReply, error) {
	return &pb.ListFundReply{}, nil
}
