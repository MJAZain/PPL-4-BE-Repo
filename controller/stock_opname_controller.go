package controller

import (
	"go-gin-auth/dto"
	"go-gin-auth/mapper"
	"go-gin-auth/service"
	"go-gin-auth/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockOpnameController struct {
	service service.StockOpnameService
}

func NewStockOpnameController(s service.StockOpnameService) *StockOpnameController {
	return &StockOpnameController{s}
}

func (c *StockOpnameController) Create(ctx *gin.Context) {

	var input dto.StockOpnameRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Respond(ctx, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	// Pemetaan dto ke model menggunakan mapper
	stockOpname := mapper.ToModelStockOpname(input)

	if err := c.service.Create(&stockOpname); err != nil {
		utils.Respond(ctx, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	utils.Respond(ctx, http.StatusCreated, "Success", nil, input)
}

func (c *StockOpnameController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.Respond(ctx, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	utils.Respond(ctx, http.StatusOK, "Success", nil, data)
}

func (c *StockOpnameController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := c.service.GetByID(uint(id))
	if err != nil {
		utils.Respond(ctx, http.StatusNotFound, "error", err.Error(), nil)
		return
	}
	utils.Respond(ctx, http.StatusOK, "Success", nil, data)
}

func (c *StockOpnameController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Cek apakah data ada terlebih dahulu
	exists, err := c.service.IsExist(uint(id))
	if err != nil {
		utils.Respond(ctx, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	if !exists {
		utils.Respond(ctx, http.StatusNotFound, "error", "Data tidak ditemukan", nil)
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		utils.Respond(ctx, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	utils.Respond(ctx, http.StatusOK, "Success", nil, "Data berhasil dihapus")
}
