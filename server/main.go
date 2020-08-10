package main

import (
	"github.com/encryptgig/scraper-grpc/scrapper"
	pb "github.com/encryptgig/scraper-grpc/scrapperpb"
	"google.golang.org/grpc"
	"log"
	"net"
)



func main() {

	log.Println("Starting Greet Server...")

	lis,err := net.Listen("tcp",":8765")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterScrapperServiceServer(s, &scrapper.Service{})
	pb.RegisterEncryptionServiceServer(s, &scrapper.Encryptor{})

	if err := s.Serve(lis) ; err != nil {
		log.Fatal("Cannot serve",err)
	}
}
