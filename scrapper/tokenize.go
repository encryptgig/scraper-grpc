package scrapper

import (
	"context"
	"github.com/encryptgig/scraper-grpc/scrapperpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Encryptor struct {

}

func (e *Encryptor) Tokenize(ctx context.Context, in *scrapperpb.TokenizationRequest) (*scrapperpb.TokenizationResponse, error) {

	if len(in.KID) == 0 {
		in.KID = "default_key"
	}

	key, ok := Keys.Get(in.KID)

	if !ok {
		return nil,status.Errorf(codes.InvalidArgument , "Invalid key id for tokenization")
	}

	out, err := Tokenize(in.Data, key.KeyByte, key.IV)
	if err != nil {
		return nil,status.Errorf(codes.Internal , "failed to tokenize")
	}

	return &scrapperpb.TokenizationResponse{KID:in.KID,Data:out},nil
}
func (e *Encryptor) DeTokenize(ctx context.Context, in *scrapperpb.TokenizationRequest) (*scrapperpb.TokenizationResponse, error) {
	if len(in.KID) == 0 {
		in.KID = "default_key"
	}

	key, ok := Keys.Get(in.KID)

	if !ok {
		return nil,status.Errorf(codes.InvalidArgument , "invalid key id for detokenization")
	}

	out, err := Detokenize(in.Data, key.KeyByte, key.IV)
	if err != nil {
		return nil,status.Errorf(codes.Internal , "failed to detokenize")
	}

	return &scrapperpb.TokenizationResponse{KID:in.KID,Data:out},nil
}

