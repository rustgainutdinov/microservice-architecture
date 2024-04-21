package usecase

import (
	"usercrud/src/domain"
)

type UserRepository interface {
	Store(domain.User) error
	Select() ([]domain.User, error)
	Find(id uint) (domain.User, error)
	Delete(id string) error
}
