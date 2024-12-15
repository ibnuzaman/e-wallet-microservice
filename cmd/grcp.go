package cmd

import (
	"ewallet-framework/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to open grpc port: ", err)
	}

	s := grpc.NewServer()

	//list method
	//pb.ExampleMethod(s, &grpc....)

	logrus.Info("GRPC Server running on port: ", helpers.GetEnv("GRPC_PORT", "7000"))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}
}
