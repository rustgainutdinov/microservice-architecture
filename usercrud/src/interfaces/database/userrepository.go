package database

import (
	"usercrud/src/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) error {
	return db.Save(&u)
}

func (db *UserRepository) Select() ([]domain.User, error) {
	var user []domain.User
	err := db.FindAll(&user)
	return user, err
}

func (db *UserRepository) Find(id uint) (domain.User, error) {
	user := domain.User{}
	err := db.FindById(&user, id)
	return user, err
}

func (db *UserRepository) Delete(id string) error {
	var user []domain.User
	return db.DeleteById(&user, id)
}
