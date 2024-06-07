package db

import (
	"demo-hex-go/internal/core/domain/entity"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresRepository struct {
	*gorm.DB
}

func NewPostgresRepository() *PostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "password"
	dbname := "demo-go"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})

	return &PostgresRepository{
		db,
	}
}
