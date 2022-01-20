package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Isomiddinov7/RestApi/modules/book_comments/controllers"
	"github.com/Isomiddinov7/RestApi/modules/book_comments/models"
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

	var book_comment models.PostBookComments

	json.Unmarshal(body, &book_comment)
	
	encoding.Encode(controllers.Post(book_comment))
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	
	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book_comment models.DeleteBookComments

	json.Unmarshal(body, &book_comment)
	
	encoding.Encode(controllers.Delete(book_comment))
}

func PUT(w http.ResponseWriter, r *http.Request) {
	
	encoding := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var book_comment models.PutBookComments

	json.Unmarshal(body, &book_comment)

	encoding.Encode(controllers.Update(book_comment))
}
