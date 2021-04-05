package Services

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jean0206/customer-api-golang/database"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	query := `{
		transactions(func:has(idTransaction), first: 200) {
			buyerId
			idTransaction
			ip
			device
			products
		}
	  }`
	transactions := database.GetDBConsult(query)

	w.Write([]byte(fmt.Sprintf("%s", string(transactions))))
}

func GetCustomerTransactionById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	query := fmt.Sprintf(`{
		customer(func:has(buyerId)) @filter(eq(buyerId,%s)){
			products
			ip	
			buyerId
	  }
	}`, userId)

	customers := database.GetDBConsult(query)

	w.Write([]byte(fmt.Sprintf("%s", string(customers))))
}
