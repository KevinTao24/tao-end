package service

import (
	"context"

	pb "tao-end/api/llm/v1"
)

type LlmService struct {
	pb.UnimplementedLlmServer
}

func NewLlmService() *LlmService {
	return &LlmService{}
}

func (s *LlmService) CreateLlm(ctx context.Context, req *pb.CreateLlmRequest) (*pb.CreateLlmReply, error) {
	return &pb.CreateLlmReply{}, nil
}
func (s *LlmService) UpdateLlm(ctx context.Context, req *pb.UpdateLlmRequest) (*pb.UpdateLlmReply, error) {
	return &pb.UpdateLlmReply{}, nil
}
func (s *LlmService) DeleteLlm(ctx context.Context, req *pb.DeleteLlmRequest) (*pb.DeleteLlmReply, error) {
	return &pb.DeleteLlmReply{}, nil
}
func (s *LlmService) GetLlm(ctx context.Context, req *pb.GetLlmRequest) (*pb.GetLlmReply, error) {
	return &pb.GetLlmReply{}, nil
}
func (s *LlmService) ListLlm(ctx context.Context, req *pb.ListLlmRequest) (*pb.ListLlmReply, error) {
	return &pb.ListLlmReply{}, nil
}
