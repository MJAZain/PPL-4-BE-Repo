package dto

type StockOpnameDetailRequest struct {
	ObatID    uint `json:"obat_id" binding:"required"`
	StokFisik int  `json:"stok_fisik" binding:"required"`
}

type StockOpnameRequest struct {
	UserID  uint                       `json:"user_id" binding:"required"`
	Details []StockOpnameDetailRequest `json:"details" binding:"required,dive"`
}
