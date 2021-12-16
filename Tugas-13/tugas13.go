package main

//ini ga berjalan sama sekali hehe
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string `json:"Nama", "Matakuliah", "IndeksNilai"`
	Nilai, ID                     uint   `json:"Nilai", "ID"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

//Post
func PostMahasiswa(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var newNilai = NilaiMahasiswa{}

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			//by body json
			if err := decodeJSON.Decode(&nilaiNilaiMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			//by form
			getId := r.PostFormValue("ID")
			id, _ := strconv.Atoi(getId)
			fmt.Println(id)
			getNama := r.PostFormValue("Nama")
			nama := getNama
			fmt.Println(nama)
			getMatkul := r.PostFormValue("Matakuliah")
			matkul := getMatkul
			fmt.Println(matkul)
			getIndeks := r.PostFormValue("IndeksNilai")
			Indeks := getIndeks
			fmt.Println(Indeks)
			getNilai := r.PostFormValue("Nilai")
			nilai, _ := strconv.Atoi(getNilai)
			fmt.Println(nilai)

			newNilai.Nama = nama
			newNilai.MataKuliah = matkul

		}

		if newNilai.Nilai >= 80 {
			newNilai.IndeksNilai = "Indeksnya A "
		} else if newNilai.Nilai >= 70 {
			newNilai.IndeksNilai = "Indeksnya B "
		} else if newNilai.Nilai >= 60 {
			newNilai.IndeksNilai = "Indeksnya C "
		} else if newNilai.Nilai >= 50 {
			newNilai.IndeksNilai = "Indeksnya D "
		} else if newNilai.Nilai < 50 {
			newNilai.IndeksNilai = "Indeksnya E "
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
		dataNilai, _ := json.Marshal(newNilai)
		rw.Write(dataNilai)
		return

	}

	http.Error(rw, "NOT ALLOWED", http.StatusMethodNotAllowed)
	return

}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
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

func main() {
	//method
	http.HandleFunc("/get_mahasiswa", getMahasiswa)
	http.HandleFunc("/post_nilai", PostMahasiswa)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
