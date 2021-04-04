package main

import (
	"net/http"

	"github.com/jean0206/customer-api-golang/routes"
)

func main() {
	/*customer := example.Customer{
		Uid:  "0x6",
		Name: "Jean Carlos",
		Age:  20,
	}*/
	http.ListenAndServe(":3000", routes.GetRoutes())
}
