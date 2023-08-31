package netxd_dal_services

import (
	"context"

	"github.com/20-VIGNESH_K/netxd-grpc/netxd-dal/netxd_dal_interfaces"
	"github.com/20-VIGNESH_K/netxd-grpc/netxd-dal/netxd_dal_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NetxdCustomerService struct {
	NetxdCustomerCollection *mongo.Collection
	ctx                     context.Context
}

func NewNetxdCustomerServiceInit(collection *mongo.Collection, ctx context.Context) netxd_dal_interfaces.INetxdCustomer {
	return &NetxdCustomerService{collection, ctx}
}

func (c *NetxdCustomerService) CreateCustomer(user *netxd_dal_models.NetxdCustomer) (*netxd_dal_models.DBResponse, error) {
	res, err := c.NetxdCustomerCollection.InsertOne(c.ctx, &user)

	if err != nil {
		return nil, err
	}

	var newUser *netxd_dal_models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.NetxdCustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}
