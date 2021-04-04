package example

import (
	"github.com/jean0206/customer-api-golang/models"
)

type Customer struct {
	Uid  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func GetHola() string {
	customer := models.Customer{
		Uid:  "0x6",
		Name: "Jean Carlos",
		Age:  20,
	}
	return customer.Name
}
