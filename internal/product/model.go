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
	Code            string            `gorm:"type:varchar(50);not null;comment:Kode/SKU" json:"code" form:"code"`
	Barcode         string            `gorm:"type:varchar(100);not null;comment:Barcode" json:"barcode" form:"barcode"`
	CategoryID      uint              `gorm:"not null;comment:ID Kategori" json:"category_id" form:"category_id"`
	Category        category.Category `gorm:"-" json:"category"`
	UnitID          uint              `gorm:"not null;comment:ID Satuan" json:"unit_id" form:"unit_id"`
	Unit            unit.Unit         `gorm:"-" json:"unit"`
	PackageContent  int               `gorm:"type:integer;not null;default:0;comment:Isi Perkemasan" json:"package_content" form:"package_content"`
	PurchasePrice   float64           `gorm:"type:decimal(15,2);not null;comment:Harga Beli" json:"purchase_price" form:"purchase_price"`
	SellingPrice    float64           `gorm:"type:decimal(15,2);not null;comment:Harga Jual" json:"selling_price" form:"selling_price"`
	WholesalePrice  float64           `gorm:"type:decimal(15,2);not null;comment:Harga Grosir" json:"wholesale_price" form:"wholesale_price"`
	StockBuffer     int               `gorm:"not null;default:0;comment:Buffer Stok" json:"stock_buffer" form:"stock_buffer"`
	StorageLocation string            `gorm:"type:varchar(100);comment:Lokasi Penyimpanan" json:"storage_location" form:"storage_location"`
	Brand           string            `gorm:"type:varchar(100);comment:Merk Dagang" json:"brand" form:"brand"`
	CreatedBy       uint              `gorm:"not null;comment:ID Pengguna Pembuat" json:"created_by"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedBy       uint              `gorm:"not null;comment:ID Pengguna Pengubah" json:"updated_by"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedBy       uint              `gorm:"comment:ID Pengguna Penghapus" json:"deleted_by"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}
