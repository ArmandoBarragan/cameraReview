package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"name"`
	Brand       string `json:"brand" gorm:"brand"`
	Description string `json:"description" gorm:"description"`
}
