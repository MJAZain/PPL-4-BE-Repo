package product

// import (
// 	"go-gin-auth/model"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type ProductHandler struct {
// 	DB             *gorm.DB
// 	ProductService *ProductService
// }

// func NewProductHandler(db *gorm.DB) *ProductHandler {
// 	return &ProductHandler{
// 		DB:             db,
// 		ProductService: NewProductService(),
// 	}
// }

// func (h *ProductHandler) CreateProduct(c *gin.Context) (*gin.H, error) {
// 	var input Product

// 	if err := c.ShouldBind(&input); err != nil {
// 		return &gin.H{
// 			"status":  "failed",
// 			"message": err.Error(),
// 		}, nil
// 	}

// 	if input.Name == "" ||
// 		input.Code == "" ||
// 		input.Barcode == "" ||
// 		input.Category == "" ||
// 		input.Unit == "" ||
// 		input.PackageContent == 0 ||
// 		input.PurchasePrice == 0 ||
// 		input.SellingPrice == 0 ||
// 		input.StockQuantity == 0 ||
// 		input.StorageLocation == "" ||
// 		input.ExpiryDate.IsZero() ||
// 		input.Brand == "" {
// 		return &gin.H{
// 			"status":  "failed",
// 			"message": "All fields are required",
// 		}, nil
// 	}
// 	currentUser, ok := c.Get("currentUser")
// 	if !ok {
// 		return &gin.H{
// 			"status":  "failed",
// 			"message": "Failed to get current user",
// 		}, nil
// 	}

// 	if currentUser == nil {
// 		return &gin.H{
// 			"status":  "failed",
// 			"message": "Failed to get current user",
// 		}, nil
// 	}

// 	input.CreatedBy = currentUser.(model.User).ID

// 	isSuccess, message := h.ProductService.CreateDataProduct(input)

// 	if isSuccess {
// 		return &gin.H{
// 			"status":  "success",
// 			"message": message,
// 		}, nil
// 	}

// 	return &gin.H{
// 		"status":  "failed",
// 		"message": message,
// 	}, nil
// }
