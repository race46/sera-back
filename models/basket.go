package models

type Basket struct {
	Username  string  `json:"username"`
	ProductId int64   `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId; constraint:OnUpdate:CASCADE,OnDelete:DELETE;"`
}
