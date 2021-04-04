package models

type Product struct {
	UID         string  `json:"uid,omitempty"`
	IdProduct   string  `json:"idProduct,omitempty"`
	NameProduct string  `json:"nameProduct,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
