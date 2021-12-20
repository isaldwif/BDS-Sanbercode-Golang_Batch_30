package service

import (
	"Tugas-1/category"
	"Tugas-1/models"
	"Tugas-1/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Post
func PostCategory(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	var newCategory models.Category

	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := category.Insert(ctx, newCategory); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)

}

//fungsiGET
func GetAllCategory(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	category, err := category.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, category, http.StatusOK)

}

// Update
func UpdateCategory(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var newCategory models.Category

	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idCategory = ps.ByName("id")

	if err := category.Update(ctx, newCategory, idCategory); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)
}

// Delete
func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idCategory = ps.ByName("id")
	if err := category.Delete(ctx, idCategory); err != nil {
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
