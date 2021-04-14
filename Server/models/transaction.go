package models

type Transaction struct {
	UID           string   `json:"uid,omitempty"`
	IdTransaction string   `json:"idTransaction,omitempty"`
	BuyerId       string   `json:"buyerId,omitempty"`
	Ip            string   `json:"ip,omitempty"`
	Device        string   `json:"device,omitempty"`
	Products      []string `json:"products,omitempty"`
}

type Transactions struct {
	transactions []Transaction `json:"transactions"`
}
