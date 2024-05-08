package handler

import (
	"encoding/json"
	"net/http"
	"shope/modle"
	"strconv"

	"github.com/gorilla/mux"
)

var Allcustomers = []modle.Costumer{
	{Id: 1, Name: "Pawan", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
	{Id: 2, Name: "Randeep", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
	{Id: 3, Name: "Ashu", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
	{Id: 4, Name: "Pardeep", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
	{Id: 5, Name: "Rahul", Mobile: "7988323110", Address: modle.Address{HouseNo: "43", Town: "kheri", District: "Rohtak"}},
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer modle.Costumer

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
	}
	Allcustomers = append(Allcustomers, customer)

}
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Allcustomers)
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

	for _, v := range Allcustomers {
		if v.Id == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}
