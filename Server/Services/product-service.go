package Services

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jean0206/customer-api-golang/database"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	query := `{
		products(func:has(nameProduct)){
			nameProduct
		  idProduct
		  price
		}
	  }`
	products := database.GetDBConsult(query)

	w.Write([]byte(fmt.Sprintf("%s", string(products))))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")
	query := fmt.Sprintf(`{
		product(func:has(idProduct)) @filter(eq(idProduct,%s)){
			idProduct
			nameProduct
			price
		}
	}`, id)
	customers := database.GetDBConsult(query)
	w.Write([]byte(fmt.Sprintf("%s", string(customers))))
}
