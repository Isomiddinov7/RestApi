package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Isomiddinov7/RestApi/modules/author/controllers"
	"github.com/Isomiddinov7/RestApi/modules/author/models"
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

	var author models.PostAuthor

	json.Unmarshal(body, &author)

	encoding.Encode(controllers.Post(author))
}

func DELETE(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var author models.DeleteAuthor

	json.Unmarshal(body, &author)

	encoding.Encode(controllers.Delete(author))
}

func PUT(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var author models.UpdateAuthor

	json.Unmarshal(body, &author)

	encoding.Encode(controllers.Put(author))
}
