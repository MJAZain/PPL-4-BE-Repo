// repository/stock_opname_repository.go
package repository

import (
	"go-gin-auth/model"

	"gorm.io/gorm"
)

type StockOpnameRepository interface {
	Create(opname *model.StockOpname) error
	GetAll() ([]model.StockOpname, error)
	GetByID(id uint) (model.StockOpname, error)
	Delete(id uint) error
	IsExist(id uint) (bool, error) // <-- ubah tipe return menjadi (bool, error)
}

type stockOpnameRepository struct {
	db *gorm.DB
}

func NewStockOpnameRepository(db *gorm.DB) StockOpnameRepository {
	return &stockOpnameRepository{db}
}

func (r *stockOpnameRepository) Create(opname *model.StockOpname) error {
	return r.db.Create(opname).Error
}

func (r *stockOpnameRepository) GetAll() ([]model.StockOpname, error) {
	var opnames []model.StockOpname
	err := r.db.Preload("Details").Find(&opnames).Error
	return opnames, err
}

func (r *stockOpnameRepository) GetByID(id uint) (model.StockOpname, error) {
	var opname model.StockOpname
	err := r.db.Preload("Details").First(&opname, id).Error
	return opname, err
}

func (r *stockOpnameRepository) Delete(id uint) error {
	return r.db.Delete(&model.StockOpname{}, id).Error
}
func (r *stockOpnameRepository) IsExist(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.StockOpname{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
