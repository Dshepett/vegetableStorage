package storage

import (
	"errors"
	"fmt"
	"log"
	"vegetableShop/internal/models"
)

type CategoryRepository struct {
	storage *Storage
}

func (c *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	rows, err := c.storage.db.Query("SELECT id,name,description,parent_id from categories")
	if err != nil {
		return nil, errors.New("error occurred during getting categories")
	}
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.Id, &category.Name, &category.Description, &category.ParentID); err != nil {
			return nil, errors.New("error occurred during getting categories")
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *CategoryRepository) FindById(id int) (*models.Category, error) {
	category := models.Category{}
	query := "SELECT id, name, description, parent_id from categories where id = $1"
	err := c.storage.db.QueryRow(query, id).Scan(&category.Id, &category.Name, &category.Description, &category.ParentID)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("category with id = %d does not exist", id))
	}
	return &category, nil
}

func (c *CategoryRepository) Update(id uint, name string, description string, parentId uint) error {
	return nil
}

func (c *CategoryRepository) Delete(id uint) error {
	return nil
}

func (c *CategoryRepository) Create(category models.Category) (*models.Category, error) {
	if category.ParentID != 0 {
		var count int
		query := "SELECT count(*) from categories where id=$1"
		if err := c.storage.db.QueryRow(query, category.ParentID).Scan(&count); err != nil {
			log.Println(err)
			return nil, errors.New("error occurred during checking parent existence")
		}
		if count == 0 {
			return nil, errors.New(fmt.Sprintf("error: no parent with id = %d", category.ParentID))
		}
	}
	query := "INSERT INTO categories(name, description, parent_id) values ($1,$2,$3) RETURNING id"
	if err := c.storage.db.QueryRow(query, category.Name, category.Description, category.ParentID).Scan(&category.Id); err != nil {
		log.Println(err)
		return nil, errors.New("error occurred during adding category")
	}
	return &category, nil
}
