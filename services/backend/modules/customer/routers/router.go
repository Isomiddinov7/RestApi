package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Isomiddinov7/RestApi/modules/customer/controllers"
	"github.com/Isomiddinov7/RestApi/modules/customer/models"
)

func GET(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	encoding.Encode(controllers.Get())
}

func POST(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var customer models.PostCustomer

	json.Unmarshal(body, &customer)

	encoding.Encode(controllers.Post(customer))
}

func DELETE(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var customer models.DeleteCustomer

	json.Unmarshal(body, &customer)

	encoding.Encode(controllers.Delete(customer))
}

func PUT(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var customer models.PutCustomer

	json.Unmarshal(body, &customer)

	encoding.Encode(controllers.Put(customer))
}
