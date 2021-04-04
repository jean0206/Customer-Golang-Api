package models

type Customer struct {
	UID  string `json:"uid,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Customers struct {
	Customers []Customer `json:"customers"`
}
