package mapper

import (
	"go-gin-auth/dto"
	"go-gin-auth/model"
	"strconv"
)

// Fungsi untuk mengonversi StockOpnameRequest menjadi StockOpname
func ToModelStockOpname(request dto.StockOpnameRequest) model.StockOpname {
	var details []model.StockOpnameDetail
	// Pemetaan untuk detail
	for _, detail := range request.Details {
		details = append(details, model.StockOpnameDetail{
			ProductID:   strconv.FormatUint(uint64(detail.ObatID), 10),
			ActualStock: detail.StokFisik,
			// StokSistem: 0, // Asumsikan nilai stok sistem untuk detail, sesuaikan jika ada data lain
			// Selisih:    0, // Asumsikan nilai selisih, sesuaikan sesuai dengan kebutuhan
		})
	}

	// Pemetaan objek StockOpname utama
	return model.StockOpname{
		CreatedBy: strconv.FormatUint(uint64(request.UserID), 10),
		Details:   details,
	}
}
