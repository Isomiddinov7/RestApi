package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Isomiddinov7/RestApi/modules/book/controllers"
	"github.com/Isomiddinov7/RestApi/modules/book/models"
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

	var book models.PostBook

	json.Unmarshal(body, &book)

	encoding.Encode(controllers.Post(book))
}

func DELETE(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book models.DeleteBook

	json.Unmarshal(body, &book)

	encoding.Encode(controllers.Delete(book))
}

func PUT(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book models.PutBook

	json.Unmarshal(body, &book)

	encoding.Encode(controllers.Update(book))
}
