package main

import (
	"context"
	"fmt"

	"github.com/20-VIGNESH_K/netxd-grpc/netxd-config/config"
	netxd_constants "github.com/20-VIGNESH_K/netxd-grpc/netxd-config/constants"
	cus "github.com/20-VIGNESH_K/netxd-grpc/netxd-customer-proto"
	controllers "github.com/20-VIGNESH_K/netxd-grpc/netxd-customer-server/netxd_controller"
	services "github.com/20-VIGNESH_K/netxd-grpc/netxd-dal/netxd_dal_services"

	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
)

func initDatabase(client *mongo.Client) {

	customerCollection := config.GetCollection(client, "bankdb", "customer")
	controllers.CustomerService= services.NewNetxdCustomerServiceInit(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxd_constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	cus.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("server running on port", netxd_constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
