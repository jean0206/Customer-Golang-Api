package Services

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jean0206/customer-api-golang/database"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	query := `{
		customers(func:has(name),first: 200){
			age
			name
		  id
		}
	  }`

	customers := database.GetDBConsult(query)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%s", string(customers))))

}

func GetCustomersWithIp(w http.ResponseWriter, r *http.Request) {
	ip := chi.URLParam(r, "ip")
	query := fmt.Sprintf(`{
		customer(func:has(ip)) @filter(eq(ip,%s)){
			ip
			buyerId
		}
	}`, ip)

	customers := database.GetDBConsult(query)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(fmt.Sprintf("%s", string(customers))))
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	query := fmt.Sprintf(`{
		customer(func:has(id)) @filter(eq(id,%s)){
			id
			name
			age
		}
	}`, id)
	customers := database.GetDBConsult(query)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%s", string(customers))))
}
