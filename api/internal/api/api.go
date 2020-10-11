package api

import (
	"fmt"
	"net"
	"strconv"

	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"xorm.io/xorm"

	"plant-manager/internal/api/apiutils"
	"plant-manager/internal/api/v1/customers"
	"plant-manager/internal/api/v1/users"
	pb "plant-manager/pb/go"
)

func Run(db *xorm.Engine, port int) {

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcmw.ChainUnaryServer(
				apiutils.ModelInjector(db),
			),
		),
	)
	initRoutes(s)

	reflection.Register(s)

	fmt.Printf("Server running on port %d\n", port)
	if err = s.Serve(lis); err != nil {
		panic("failed to serve grpc: " + err.Error())
	}
}

func initRoutes(s *grpc.Server) {
	pb.RegisterV1UsersServer(s, users.InitRoutes())
	pb.RegisterV1CustomersServer(s, customers.InitRoutes())
}