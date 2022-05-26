package storage

import (
	"errors"
	"fmt"
	"log"
	"vegetableShop/internal/models"
)

type UserRepository struct {
	storage *Storage
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	rows, err := u.storage.db.Query("SELECT id,name,email,password,role from users")
	if err != nil {
		return nil, errors.New("error occurred during getting users")
	}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword, &user.Role); err != nil {
			return nil, errors.New("error occurred during getting users")
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) FindById(id uint) (*models.User, error) {
	user := models.User{}
	query := "SELECT id,name,email,password,role from users where id = $1"
	err := u.storage.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword, &user.Role)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user with id = %d does not exist", id))
	}
	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	query := "SELECT id,name,email,password,role from users where email = $1"
	err := u.storage.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword, &user.Role)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("wrong email"))
	}
	return &user, nil
}

func (u *UserRepository) Update(user models.User) error {
	query := "UPDATE users SET name = $1, password=$2, email = $3, role = &4 WHERE id = $3;"
	if _, err := u.storage.db.Exec(query, user.Name, user.HashedPassword, user.Email, user.Role, user.Id); err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(id uint) error {
	if exist, err := u.Exist(id); err != nil {
		return err
	} else if !exist {
		return errors.New(fmt.Sprintf("user with id =%d does not exist", id))
	}
	_, err := u.storage.db.Exec("DELETE FROM users WHERE id =$1", id)
	return err
}

func (u *UserRepository) Create(user models.User) (*models.User, error) {
	query := "INSERT INTO users(name, email, password, role) values ($1,$2,$3, $4) RETURNING id"
	if err := u.storage.db.QueryRow(query, user.Name, user.Email, user.HashedPassword, 1).Scan(&user.Id); err != nil {
		log.Println(err)
		return nil, errors.New("error occurred during adding user")
	}
	return &user, nil
}

func (u *UserRepository) Exist(id uint) (bool, error) {
	var count int
	if err := u.storage.db.QueryRow("SELECT COUNT(*) FROM users WHERE id=$1", id).Scan(&count); err != nil {
		return false, errors.New("error during checking user existence")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
