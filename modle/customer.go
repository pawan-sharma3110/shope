package modle

type Costumer struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Mobile  string   `json:"mobile"`
	Address Address  `json:"address"`
	Product []string `json:"product"`
}

type Address struct {
	HouseNo  string `json:"house_no"`
	Town     string `json:"town"`
	District string `json:"district"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
}
