package mapper

import (
	"go-gin-auth/dto"
	"go-gin-auth/model"
)

// Fungsi untuk mengonversi StockOpnameRequest menjadi StockOpname
func ToModelStockOpname(request dto.StockOpnameRequest) model.StockOpname {
	var details []model.StockOpnameDetail
	// Pemetaan untuk detail
	for _, detail := range request.Details {
		details = append(details, model.StockOpnameDetail{
			ObatID:    detail.ObatID,
			StokFisik: detail.StokFisik,
			// StokSistem: 0, // Asumsikan nilai stok sistem untuk detail, sesuaikan jika ada data lain
			// Selisih:    0, // Asumsikan nilai selisih, sesuaikan sesuai dengan kebutuhan
		})
	}

	// Pemetaan objek StockOpname utama
	return model.StockOpname{
		UserID:  request.UserID,
		Details: details,
	}
}
