package models

type Product struct {
	ID        uint   `gorm:"primaryKey; autoIncrement;" json:"id"`
	Username  string `json:"username"`
	Price     int64  `json:"price"`
	Type      string `gorm:"not null" json:"type"`
	Active    bool   `json:"active"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Location  string `json:"location"`
}
