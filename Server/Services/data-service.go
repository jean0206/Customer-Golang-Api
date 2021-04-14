package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"strings"

	"github.com/jean0206/customer-api-golang/database"
	"github.com/jean0206/customer-api-golang/models"
)

func ImportData(w http.ResponseWriter, r *http.Request) {

	body := r.FormValue("date")
	if body != "" {
		date, err := strconv.ParseInt(body, 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{'message':'error adding the data'}"))
		}
		log.Println(date)
		go ImportProducts(date)
		go ImportCustomers(date)
		go ImportTransactions(date)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'message':'information has been added'}"))
	} else {
		date := time.Now().Unix()
		go ImportProducts(date)
		go ImportCustomers(date)
		ImportTransactions(date)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'message':'information has been added'}"))
	}

}

func ImportCustomers(actualDate int64) string {

	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=" + strconv.FormatInt(actualDate, 16)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var customers []models.Customer
	if err := json.Unmarshal([]byte(body), &customers); err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(customers)

	if err != nil {
		fmt.Println(err)
	} else {
		database.MutateDatabase(jsonBytes)
	}

	return string("success")
}

func ImportProducts(actualDate int64) {
	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date=" + strconv.FormatInt(actualDate, 16)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	productsArray := strings.Split(string(body), "\n")

	for _, element := range productsArray {
		separateElements := strings.Split(element, "'")
		if len(separateElements) >= 3 {
			id := separateElements[0]
			name := ""
			price := separateElements[len(separateElements)-1]

			if len(separateElements) > 3 {
				for i := 1; i < len(separateElements)-1; i++ {
					name += (separateElements[i])
				}
			} else {
				name = separateElements[1]
			}

			var newPrice float64

			if s, err := strconv.ParseFloat(price, 64); err == nil {
				newPrice = s
			}

			newProduct := models.Product{"_:newId", id, name, newPrice}

			jsonBytes, err := json.Marshal(newProduct)
			if err != nil {
				fmt.Println(err)
			} else {
				database.MutateDatabase(jsonBytes)
			}

			log.Println(newProduct.NameProduct)
			log.Println(newProduct.Price)

		}
	}

}

func ImportTransactions(actualDate int64) {
	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=" + strconv.FormatInt(actualDate, 16)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(body)
	newData := strings.Split(data, ")")

	for _, row := range newData {
		reg, err := regexp.Compile("[^a-zA-Z0-9-#-,-.]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(row, "/")
		values := strings.Split(processedString, "/")
		if len(values) >= 6 {
			id := values[1]
			buyerId := values[2]
			ip := values[3]
			device := values[4]
			product := strings.Split(values[5], ",")
			product[0] = strings.Replace(product[0], "(", "", -1)
			for i := 0; i < len(product); i++ {
				product[i] = string(product[i])
			}

			newTransaction := models.Transaction{"_:newId", id, buyerId, ip, device, product}

			jsonBytes, err := json.Marshal(newTransaction)
			if err != nil {
				fmt.Println(err)
			} else {
				database.MutateDatabase(jsonBytes)
			}
		}

	}
}
