package usecase

import "usercrud/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) error {
	return i.UserRepository.Store(u)
}

func (i *UserInteractor) Update(u domain.User) error {
	_, err := i.UserRepository.Find(u.ID)
	if err != nil {
		return err
	}
	return i.UserRepository.Store(u)
}

func (i *UserInteractor) GetInfo() ([]domain.User, error) {
	return i.UserRepository.Select()
}

func (i *UserInteractor) Delete(id string) error {
	return i.UserRepository.Delete(id)
}
