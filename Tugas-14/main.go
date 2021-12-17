//main.go, digunakan untuk menjalankan melakukan aksi terhadap data, bisa di katakan sebuah controller.

package main

import (
	"Tugas-14/config"
	"Tugas-14/function"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	//controller
	router := httprouter.New()
	router.GET("/mahasiswa", function.GetMahasiswa)
	router.POST("/mahasiswa/create", function.PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", function.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", function.DeleteMahasiswa)

	//casting
	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
