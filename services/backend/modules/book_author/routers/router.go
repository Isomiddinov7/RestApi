package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Isomiddinov7/RestApi/modules/book_author/controllers"
	"github.com/Isomiddinov7/RestApi/modules/book_author/models"
)

func GET(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	encoding.Encode(controllers.Get())
}

func POST(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book_author models.PostBookAuthor

	json.Unmarshal(body, &book_author)

	encoding.Encode(controllers.Post(book_author))
}

func DELETE(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book_author models.DeleteBookAuthor

	json.Unmarshal(body, &book_author)

	encoding.Encode(controllers.Delete(book_author))
}

func PUT(w http.ResponseWriter, r *http.Request) {

	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book_author models.PutBookAuthor

	json.Unmarshal(body, &book_author)

	encoding.Encode(controllers.Update(book_author))
}
