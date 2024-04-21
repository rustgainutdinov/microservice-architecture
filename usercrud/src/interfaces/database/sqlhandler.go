package database

type SqlHandler interface {
	FindAll(object interface{}) error
	FindById(obj interface{}, id uint) error
	Save(object interface{}) error
	DeleteById(object interface{}, id string) error
}
