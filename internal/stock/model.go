package stock

import (
	"time"
)

type Stock struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null;comment:ID Produk" json:"product_id"`
	Quantity  int       `gorm:"not null;comment:Kuantitas" json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
