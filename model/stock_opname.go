package model

import (
	"time"
)

// type StockOpname struct {
// 	ID        uint                `gorm:"primaryKey" json:"id"`
// 	UserID    uint                `gorm:"not null" json:"user_id"`
// 	CreatedAt time.Time           `gorm:"autoCreateTime" json:"created_at"`
// 	Details   []StockOpnameDetail `gorm:"foreignKey:StockOpnameID" json:"details,omitempty"`
// }

type StockOpname struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	UserID    uint                `gorm:"not null" json:"user_id"`
	CreatedAt time.Time           `gorm:"autoCreateTime" json:"created_at"`
	Details   []StockOpnameDetail `gorm:"foreignKey:StockOpnameID;constraint:OnDelete:CASCADE;" json:"details,omitempty"`
}
