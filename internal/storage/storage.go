package storage

import (
	"database/sql"
	"fmt"
	"log"
	"vegetableShop/internal/config"
)

type Storage struct {
	config             *config.Config
	db                 *sql.DB
	userRepository     *UserRepository
	categoryRepository *CategoryRepository
}

func New(config *config.Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open() {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		s.config.DBUser, s.config.DBPassword, s.config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.db = db
	s.userRepository = &UserRepository{
		storage: s,
	}
	s.categoryRepository = &CategoryRepository{
		storage: s,
	}
}

func (s *Storage) User() *UserRepository {
	return s.userRepository
}

func (s *Storage) Category() *CategoryRepository {
	return s.categoryRepository
}
