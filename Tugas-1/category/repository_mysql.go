//movie/repository_mysql.go, digunakan untuk melakukan query ke database.

package category

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
	table          = "category"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Category, error) {
	var categorys []models.Category
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
		var category models.Category
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&category.ID,
			&category.Name,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		categorys = append(categorys, category)
	}
	return categorys, nil
}

// Insert
func Insert(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id, name, created_at, updated_at) values(%v,'%v', NOW(), NOW())", table,
		category.ID,
		category.Name)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil

}

// Update
func Update(ctx context.Context, category models.Category, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name='%v' where id=%v", table,
		category.Name,
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
