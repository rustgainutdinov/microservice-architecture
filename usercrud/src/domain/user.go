package domain

type User struct {
	ID    uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  *string `json:"name" gorm:"not null"`
	Login string  `json:"login" gorm:"unique"`
}
