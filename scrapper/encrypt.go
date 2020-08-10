package scrapper

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/encryptgig/scraper-grpc/scrapperpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// pkcs5Padding is a pkcs5 padding struct.
type pkcs5 struct{}

var (
	ErrPaddingSize = errors.New("padding size error")
)
var (
	PKCS5 = &pkcs5{}
)
var (
	// difference with pkcs5 only block must be 8
	PKCS7 = &pkcs5{}
)

// Padding implements the Padding interface Padding method.
func (p *pkcs5) Padding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := blockSize - (srcLen % blockSize)
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(src, padText...)
}

// Unpadding implements the Padding interface Unpadding method.
func (p *pkcs5) Unpadding(src []byte, blockSize int) ([]byte, error) {
	srcLen := len(src)
	paddingLen := int(src[srcLen-1])
	if paddingLen >= srcLen || paddingLen > blockSize {
		return nil, ErrPaddingSize
	}
	return src[:srcLen-paddingLen], nil
}

func (e *Encryptor) Encrypt(ctx context.Context, in *scrapperpb.EncryptionRequest) (*scrapperpb.EncryptionResponse, error) {
	if len(in.KID) == 0 {
		in.KID = "default_key"
	}

	key, ok := Keys.Get(in.KID)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid key id for encryption")
	}

	block, err := aes.NewCipher(key.KeyByte)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid key bytes for encryption")
	}

	msg := PKCS7.Padding([]byte(in.Data), block.BlockSize())

	if len(msg) < block.BlockSize() || len(msg)%block.BlockSize() != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "data length error on encryption")
	}

	ciphertext := make([]byte, len(msg))

	cipher.NewCBCEncrypter(block, key.IV).CryptBlocks(ciphertext, msg)

	return &scrapperpb.EncryptionResponse{KID: in.KID, Data: base64.StdEncoding.EncodeToString(ciphertext)}, nil
}

func (e *Encryptor) Decrypt(ctx context.Context, in *scrapperpb.EncryptionRequest) (*scrapperpb.EncryptionResponse, error) {
	if len(in.KID) == 0 {
		in.KID = "default_key"
	}

	key, ok := Keys.Get(in.KID)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid key id for decryption")
	}

	block, err := aes.NewCipher(key.KeyByte)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid key bytes for decryption")
	}

	data, err := base64.StdEncoding.DecodeString(in.Data)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid data for decryption")
	}

	if len(data) < block.BlockSize() || len(data)%block.BlockSize() != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid data length for decryption")
	}

	ciphertext := make([]byte, len(data))

	cipher.NewCBCDecrypter(block, key.IV).CryptBlocks(ciphertext, data)

	msg, err := PKCS7.Unpadding(ciphertext, block.BlockSize())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid decrypted data")
	}
	return &scrapperpb.EncryptionResponse{KID: in.KID, Data: string(msg)}, nil
}
