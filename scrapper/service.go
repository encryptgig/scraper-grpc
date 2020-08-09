package scrapper

import (
	"context"
	"github.com/encryptgig/scraper-grpc/scrapperpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
}

func (s *Service) CreateKey(ctx context.Context, k *scrapperpb.KeyCreateRequest) (*scrapperpb.KeyCreateResponse, error) {
	id,err := ConfigureKey(k.Key)
	if err != nil {
		return nil,status.Errorf(codes.InvalidArgument , err.Error())
	}
	res := &scrapperpb.KeyCreateResponse{ID:id}
	return res,nil
}

func (s *Service) ListKeys(context.Context, *scrapperpb.Empty) (*scrapperpb.KeyList, error) {
	return ListKeys(), nil
}

func (s *Service) GetKeyInfo(context.Context, *scrapperpb.KeyInfoRequest) (*scrapperpb.KeyInfoResponse, error) {
	return nil, nil
}

func (s *Service) ExportKey(context.Context, *scrapperpb.KeyExportRequest) (*scrapperpb.KeyExportResponse, error) {
	return nil,nil
}

func (s *Service) DeleteKey(ctx context.Context, in *scrapperpb.KeyDeleteRequest) (*scrapperpb.KeyDeleteResponse, error) {
	return &scrapperpb.KeyDeleteResponse{ID:DeleteKey(in.ID)}, nil
}
