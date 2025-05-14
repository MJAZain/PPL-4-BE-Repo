package incomingProducts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandlerIncomingProducts() *Handler {
	return &Handler{service: NewService()}
}

// CreateIncomingProduct godoc
// @Summary Membuat produk masuk baru
// @Description Membuat produk masuk beserta detailnya
// @Tags IncomingProducts
// @Accept json
// @Produce json
// @Param incomingProduct body IncomingProductRequest true "Data produk masuk"
// @Success 201 {object} IncomingProduct
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products [post]
func (h *Handler) CreateIncomingProduct(c *gin.Context) {
	var request struct {
		IncomingProduct IncomingProduct         `json:"incoming_product"`
		Details         []IncomingProductDetail `json:"details"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.CreateIncomingProduct(&request.IncomingProduct, request.Details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal membuat produk masuk",
			"error":   err.Error(),
		})
		return
	}

	// Get full data with ID
	product, err := h.service.GetIncomingProductByID(request.IncomingProduct.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Produk masuk berhasil dibuat tetapi gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	details, err := h.service.GetIncomingProductDetails(product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Produk masuk berhasil dibuat tetapi gagal mengambil detail",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Produk masuk berhasil dibuat",
		"data": gin.H{
			"incoming_product": product,
			"details":          details,
		},
	})
}

// GetAllIncomingProducts godoc
// @Summary Mendapatkan semua produk masuk
// @Description Mendapatkan daftar semua produk masuk
// @Tags IncomingProducts
// @Produce json
// @Success 200 {array} IncomingProduct
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products [get]
func (h *Handler) GetAllIncomingProducts(c *gin.Context) {
	products, err := h.service.GetAllIncomingProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mendapatkan daftar produk masuk",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   products,
	})
}

// GetIncomingProductByID godoc
// @Summary Mendapatkan produk masuk berdasarkan ID
// @Description Mendapatkan detail produk masuk berdasarkan ID
// @Tags IncomingProducts
// @Produce json
// @Param id path int true "ID Produk Masuk"
// @Success 200 {object} IncomingProduct
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products/{id} [get]
func (h *Handler) GetIncomingProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter ID tidak valid",
			"error":   err.Error(),
		})
		return
	}

	product, err := h.service.GetIncomingProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Produk masuk tidak ditemukan",
			"error":   err.Error(),
		})
		return
	}

	details, err := h.service.GetIncomingProductDetails(product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mendapatkan detail produk masuk",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"incoming_product": product,
			"details":          details,
		},
	})
}

// UpdateIncomingProduct godoc
// @Summary Mengupdate produk masuk
// @Description Mengupdate data produk masuk berdasarkan ID
// @Tags IncomingProducts
// @Accept json
// @Produce json
// @Param id path int true "ID Produk Masuk"
// @Param incomingProduct body IncomingProduct true "Data produk masuk"
// @Success 200 {object} IncomingProduct
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products/{id} [put]
func (h *Handler) UpdateIncomingProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter ID tidak valid",
			"error":   err.Error(),
		})
		return
	}

	var incomingProduct IncomingProduct
	if err := c.ShouldBindJSON(&incomingProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.UpdateIncomingProduct(uint(id), &incomingProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengupdate produk masuk",
			"error":   err.Error(),
		})
		return
	}

	product, err := h.service.GetIncomingProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Produk masuk berhasil diupdate tetapi gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Produk masuk berhasil diupdate",
		"data":    product,
	})
}

// UpdateIncomingProductDetails godoc
// @Summary Mengupdate detail produk masuk
// @Description Mengupdate detail produk masuk berdasarkan ID produk masuk
// @Tags IncomingProducts
// @Accept json
// @Produce json
// @Param id path int true "ID Produk Masuk"
// @Param details body []IncomingProductDetail true "Detail produk masuk"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products/{id}/details [put]
func (h *Handler) UpdateIncomingProductDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter ID tidak valid",
			"error":   err.Error(),
		})
		return
	}

	var details []IncomingProductDetail
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// Pastikan semua detail memiliki incoming_product_id yang sama
	for i := range details {
		details[i].IncomingProductID = uint(id)
	}

	if err := h.service.UpdateIncomingProductDetails(details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengupdate detail produk masuk",
			"error":   err.Error(),
		})
		return
	}

	updatedDetails, err := h.service.GetIncomingProductDetails(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Detail produk masuk berhasil diupdate tetapi gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	// Update total amount pada produk masuk
	product, err := h.service.GetIncomingProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Detail produk masuk berhasil diupdate tetapi gagal mengambil produk masuk",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Detail produk masuk berhasil diupdate",
		"data": gin.H{
			"incoming_product": product,
			"details":          updatedDetails,
		},
	})
}

// DeleteIncomingProduct godoc
// @Summary Menghapus produk masuk
// @Description Menghapus produk masuk berdasarkan ID
// @Tags IncomingProducts
// @Produce json
// @Param id path int true "ID Produk Masuk"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /incoming-products/{id} [delete]
func (h *Handler) DeleteIncomingProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter ID tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// Pastikan produk masuk ada
	_, err = h.service.GetIncomingProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Produk masuk tidak ditemukan",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.DeleteIncomingProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghapus produk masuk",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Produk masuk berhasil dihapus",
	})
}
