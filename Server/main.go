package main

import (
	"net/http"

	"github.com/jean0206/customer-api-golang/routes"
)

func main() {

	http.ListenAndServe(":3000", routes.GetRoutes())
}
