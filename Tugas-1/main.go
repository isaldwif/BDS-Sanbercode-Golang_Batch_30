//main.go, digunakan untuk menjalankan melakukan aksi terhadap data, bisa di katakan sebuah controller.

package main

import (
	"Tugas-1/config"
	"Tugas-1/service"
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

	//Book
	router.GET("/books", service.GetAllBook)
	router.POST("/books", service.PostBook)
	router.PUT("/books/:id", service.UpdateBook)
	router.DELETE("/books/:id", service.DeleteBook)
	//router.GET("/categories/:id/books", service.GetIdCategory)

	//Category
	router.GET("/categories", service.GetAllCategory)
	router.POST("/categories", service.PostCategory)
	router.PUT("/categories/:id", service.UpdateCategory)
	router.DELETE("/categories/:id", service.DeleteCategory)

	//casting
	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
