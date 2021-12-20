//movie/repository_mysql.go, digunakan untuk melakukan query ke database.

package book

import (
	"Tugas-1/config"
	"Tugas-1/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "book"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.CategoryId,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}
	return books, nil
}

// Insert
func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, category_id ,created_at, updated_at) values('%v',' %v', '%v', %v, '%v', %v, '%v', %v, NOW(), NOW())", table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness,
		book.CategoryId)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil

}

// Update
func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title='%v', description='%v', image_url='%v', release_year=%v, price='%v', total_page=%v, thickness='%v', category_id=%v where id=%v", table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness,
		book.CategoryId,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
