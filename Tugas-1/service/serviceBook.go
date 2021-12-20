package service

import (
	"Tugas-1/book"
	"Tugas-1/models"
	"Tugas-1/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//fungsi book
func GetThickness(total_page int) string {
	switch {
	case total_page <= 100:
		return "tipis"
	case total_page >= 101 && total_page <= 200:
		return "sedang"
	case total_page >= 201:
		return "tebal"
	default:
		return "tipis"
	}
}

//Post
func PostBook(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if newBook.Release_year <= 1980 || newBook.Release_year >= 2021 {
		http.Error(rw, "MINIMAL TAHUN ADALAH 1980, DAN MAKSIMAL 2021", http.StatusBadRequest)
		return
	}
	newBook.Thickness = GetThickness(newBook.Total_page)

	if err := book.Insert(ctx, newBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)

}

//fungsiGET
func GetAllBook(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	book, err := book.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, book, http.StatusOK)

}

// Update
func UpdateBook(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idBook = ps.ByName("id")

	if err := book.Update(ctx, newBook, idBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)
}

// Delete
func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idBook = ps.ByName("id")
	if err := book.Delete(ctx, idBook); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}
