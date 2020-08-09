package main

import (
	"context"
	"fmt"
	"github.com/encryptgig/scraper-grpc/scrapperpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	c,err := grpc.Dial("localhost:8765",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	cl := scrapperpb.NewScrapperServiceClient(c)

	list,err := cl.ListKeys(context.Background(),&scrapperpb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*list)

	createReq := scrapperpb.KeyCreateRequest{Key:&scrapperpb.Key{KeyAlgorithm:"aes",KeySize:256,KeyIv:"F162046AFC1F2302FE7D11E4629068F1",KeyBytes:"C1AB55FE70ACC6EEB61FCDB0DF48FF799C35E1C1B3835AB86203FF8582734D09"}}
	_,err = cl.CreateKey(context.Background() , &createReq)
	if err != nil {
		log.Fatal(err)
	}
	_,err = cl.CreateKey(context.Background() , &createReq)
	if err != nil {
		log.Fatal(err)
	}
	list,err = cl.ListKeys(context.Background(),&scrapperpb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*list)
}
