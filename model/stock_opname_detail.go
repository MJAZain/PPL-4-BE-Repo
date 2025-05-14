package model

import (
	"time"
)

type StockOpnameDetail struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	StockOpnameID uint        `gorm:"not null" json:"stock_opname_id"`
	ObatID        uint        `gorm:"not null" json:"obat_id"`
	StokSistem    int         `gorm:"not null" json:"stok_sistem"`
	StokFisik     int         `gorm:"not null" json:"stok_fisik"`
	Selisih       int         `gorm:"not null" json:"selisih"`
	StockOpname   StockOpname `gorm:"foreignKey:StockOpnameID" json:"stock_opname,omitempty"`
}
