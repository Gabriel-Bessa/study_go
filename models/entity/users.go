package entity

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
