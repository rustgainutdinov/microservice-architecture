package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler(DBHost string, DBPort int, DBName string, DBUserName string, DBPass string) *SqlHandler {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUserName, DBPass, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) FindAll(obj interface{}) error {
	return handler.db.Find(obj).Error
}

func (handler *SqlHandler) FindById(obj interface{}, id uint) error {
	return handler.db.First(obj, id).Error
}

func (handler *SqlHandler) Save(obj interface{}) error {
	return handler.db.Save(obj).Error
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) error {
	return handler.db.Delete(obj, id).Error
}

func (handler *SqlHandler) Migrate(obj interface{}) error {
	return handler.db.AutoMigrate(obj)
}
