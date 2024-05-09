package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shope/db"
	"shope/modle"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	Id      int64  `json:"id,omitempty"`
	Massage string `json:"massage,omitempty"`
}

// var Allcustomers = []modle.Customer{
// 	{Id: 1, Name: "Pawan", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
// 	{Id: 2, Name: "Randeep", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
// 	{Id: 3, Name: "Ashu", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
// 	{Id: 4, Name: "Pardeep", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
// 	{Id: 5, Name: "Rahul", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
// }

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer modle.Customer

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}
	inseartID := db.InsertCustomer(customer)
	res := response{
		Id:      inseartID,
		Massage: "New Customer created",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}

}
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := db.GetAllCustomers()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}

}
func SearchById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsId, ok := params["id"]
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID parameter not provided"})
		return
	}

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	customer, err := db.GetCustomerById(int64(id))
	if err != nil {
		log.Fatal(err)

	}

	json.NewEncoder(w).Encode(customer)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	paramsId := mux.Vars(r)
	id, err := strconv.Atoi(paramsId["id"])
	if err != nil {
		log.Fatal(err)
	}
	var customer modle.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Fatal(err)
	}
	updatedRows := db.UpdateById(int64(id), customer)
	msg := fmt.Sprintf("Customer uodated successfully %v", updatedRows)
	res := response{
		Id:      int64(id),
		Massage: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	deleteRow := db.DeleteCustomerById(int64(id))
	msg := fmt.Sprintf("Customer Delete from records %v", deleteRow)
	res := response{
		Id:      int64(id),
		Massage: msg,
	}
	json.NewEncoder(w).Encode(res)
}
