package netxd_dal_interfaces

import "github.com/20-VIGNESH_K/netxd-grpc/netxd-dal/netxd_dal_models"

type INetxdCustomer interface {
	CreateCustomer(customer *netxd_dal_models.NetxdCustomer) (*netxd_dal_models.DBResponse, error)
}
