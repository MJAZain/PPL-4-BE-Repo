package product

import (
	"go-gin-auth/internal/category"
	"go-gin-auth/internal/unit"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID              uint              `gorm:"primaryKey" json:"id"`
	Name            string            `gorm:"type:varchar(255);not null;comment:Nama Obat" json:"name" form:"name"`
	Code            string            `gorm:"type:varchar(50);not null;uniqueIndex;comment:Kode/SKU" json:"code" form:"code"`
	Barcode         string            `gorm:"type:varchar(100);not null;comment:Barcode" json:"barcode" form:"barcode"`
	CategoryID      uint              `gorm:"not null;comment:ID Kategori" json:"category_id" form:"category_id"`
	Category        category.Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	UnitID          uint              `gorm:"not null;comment:ID Satuan" json:"unit_id" form:"unit_id"`
	Unit            unit.Unit         `gorm:"foreignKey:UnitID;references:ID" json:"unit"`
	PackageContent  int               `gorm:"type:integer;not null;default:0;comment:Isi Perkemasan" json:"package_content" form:"package_content"`
	PurchasePrice   float64           `gorm:"type:decimal(15,2);not null;comment:Harga Beli" json:"purchase_price" form:"purchase_price"`
	SellingPrice    float64           `gorm:"type:decimal(15,2);not null;comment:Harga Jual" json:"selling_price" form:"selling_price"`
	StockQuantity   int               `gorm:"not null;default:0;comment:Jumlah Stok" json:"stock_quantity" form:"stock_quantity"`
	StorageLocation string            `gorm:"type:varchar(100);comment:Lokasi Penyimpanan" json:"storage_location" form:"storage_location"`
	ExpiryDate      time.Time         `gorm:"type:date;comment:Tanggal Kedaluarsa" json:"expiry_date" form:"expiry_date"`
	Brand           string            `gorm:"type:varchar(100);comment:Merk Dagang" json:"brand" form:"brand"`
	CreatedBy       uint              `gorm:"not null;comment:ID Pengguna Pembuat" json:"created_by"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedBy       uint              `gorm:"not null;comment:ID Pengguna Pengubah" json:"updated_by"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedBy       uint              `gorm:"comment:ID Pengguna Penghapus" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}
