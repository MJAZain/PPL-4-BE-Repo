// DTO untuk menyesuaikan stok produk
package dto

import "time"

type StockAdjustmentRequest struct {
	ActualStock    int       `json:"actual_stock" binding:"required"`
	AdjustmentNote string    `json:"adjustment_note"`
	OpnameDate     time.Time `json:"opname_date" binding:"required"`
	PerformedBy    string    `json:"performed_by" binding:"required"`
}

// DTO untuk histori penyesuaian stok
type StockAdjustmentHistory struct {
	AdjustmentID          string    `json:"adjustment_id"`
	ProductID             string    `json:"product_id"`
	Name                  string    `json:"name"`
	PreviousStock         int       `json:"previous_stock"`
	ActualStock           int       `json:"actual_stock"`
	Discrepancy           int       `json:"discrepancy"`
	DiscrepancyPercentage float64   `json:"discrepancy_percentage"`
	AdjustmentNote        string    `json:"adjustment_note"`
	OpnameDate            time.Time `json:"opname_date"`
	PerformedBy           string    `json:"performed_by"`
}

// DTO untuk selisih stok yang signifikan
type StockDiscrepancy struct {
	ProductID             string    `json:"product_id"`
	Name                  string    `json:"name"`
	Category              string    `json:"category"`
	PreviousStock         int       `json:"previous_stock"`
	ActualStock           int       `json:"actual_stock"`
	Discrepancy           int       `json:"discrepancy"`
	DiscrepancyPercentage float64   `json:"discrepancy_percentage"`
	Flag                  string    `json:"flag"`
	OpnameDate            time.Time `json:"opname_date"`
	PerformedBy           string    `json:"performed_by"`
}

type ProductStockResponse struct {
	Name            string            `json:"name"`
	Code            string            `json:"code"`
	StockBuffer     int               `json:"stock_buffer"`
	StorageLocation string            `json:"storage_location"`
	Category        CategorySimpleDTO `json:"category"`
	Unit            UnitSimpleDTO     `json:"unit"`
}

type CategorySimpleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UnitSimpleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
