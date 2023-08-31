package netxd_dal_models

import "time"

type NetxdCustomer struct {
	CustomerId int32     `json:"customerId" bson:"customerId"`
	FirstName  string    `json:"firstName" bson:"firstName"`
	LastName   string    `json:"lastName" bson:"lastName"`
	BankId     int32     `json:"bankId" bson:"bankId"`
	Balance    float32   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	IsActive   bool      `json:"isActive" bson:"isActive"`
}

type DBResponse struct {
	CustomerId int32     `json:"customerId" bson:"customerId"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}
