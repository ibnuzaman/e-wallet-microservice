package cmd

import (
	"ewallet-framework/cmd/proto/tokenvalidation"
	"ewallet-framework/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	dependency := dependencyInject()

	s := grpc.NewServer()

	//list method
	tokenvalidation.RegisterTokenValidationServer(s, dependency.TokenValidation)
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7001"))
	if err != nil {
		log.Fatal("Failed to open grpc port: ", err)
	}

	//pb.ExampleMethod(s, &grpc....)

	logrus.Info("GRPC Server running on port: ", helpers.GetEnv("GRPC_PORT", "7001"))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}
}
