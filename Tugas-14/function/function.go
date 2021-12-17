package function

import (
	"Tugas-14/mahasiswa"
	"Tugas-14/models"
	"Tugas-14/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//punya mahasiswa

var nilaiNilaiMahasiswa = []models.NilaiMahasiswa{}

//Post
func PostMahasiswa(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var newNilai = models.NilaiMahasiswa{
			ID: uint(len(nilaiNilaiMahasiswa) + 1),
		}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			//by body json
			if err := decodeJSON.Decode(&newNilai); err != nil {
				log.Fatal(err)
			}
		} else {
			//by form
			nama := r.FormValue("nama")
			matkul := r.FormValue("matakuliah")
			getNilai := r.FormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			newNilai.Nama = nama
			newNilai.MataKuliah = matkul
			newNilai.Nilai = uint(nilai)

		}

		if newNilai.Nilai > 100 {
			http.Error(rw, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
			return
		}

		//get indeks berdasarkan nilai
		newNilai.IndeksNilai = getIndeksNilai(newNilai.Nilai)

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
		dataNilai, _ := json.Marshal(newNilai)
		rw.Write(dataNilai)
		return

	}

	http.Error(rw, "NOT FOUND", http.StatusMethodNotAllowed)
	return

}

//fungsiGET
func GetMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// Update
// UpdateMovie
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var newNilai = models.NilaiMahasiswa{}

	if err := json.NewDecoder(r.Body).Decode(&newNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMahasiswa = ps.ByName("id")

	if err := mahasiswa.Update(ctx, newNilai, idMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMovie
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMahasiswa = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMahasiswa); err != nil {
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

//fungsi indeks
func getIndeksNilai(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else if nilai < 50 {
		return "E"
	}
	return ""
}
