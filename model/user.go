package model

type User struct {
	Id       int64  `gorm:"primaryKey AUTO_INCREMENT" json:"id"`
	Email    string `gorm:"varchar(50)" json:"email"`
	Password string `gorm:"varchar (60)" json:"password"`
}