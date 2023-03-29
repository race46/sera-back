package models

type User struct {
	Username  string `gorm:"uniqueIndex;not null; primaryKey" json:"username"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
