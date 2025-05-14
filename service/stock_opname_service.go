package service

import (
	"go-gin-auth/config"
	"go-gin-auth/internal/product"
	"go-gin-auth/model"
	"go-gin-auth/repository"
)

type StockOpnameService interface {
	Create(opname *model.StockOpname) error
	GetAll() ([]model.StockOpname, error)
	GetByID(id uint) (model.StockOpname, error)
	Delete(id uint) error
	IsExist(id uint) (bool, error) // <-- ubah tipe return menjadi (bool, error)
}

type stockOpnameService struct {
	repo repository.StockOpnameRepository
}

func NewStockOpnameService(r repository.StockOpnameRepository) StockOpnameService {
	return &stockOpnameService{r}
}

func (s *stockOpnameService) Create(opname *model.StockOpname) error {
	db := config.DB
	for i := range opname.Details {
		d := &opname.Details[i]
		var obat product.Product
		if err := db.First(&obat, d.ObatID).Error; err != nil {
			return err
		}

		d.StokSistem = obat.StockBuffer
		d.Selisih = d.StokFisik - d.StokSistem

		if d.Selisih != 0 {
			obat.StockBuffer = d.StokFisik
			if err := db.Save(&obat).Error; err != nil {
				return err
			}
		}
	}
	return s.repo.Create(opname)
}

func (s *stockOpnameService) GetAll() ([]model.StockOpname, error) {
	return s.repo.GetAll()
}

func (s *stockOpnameService) GetByID(id uint) (model.StockOpname, error) {
	return s.repo.GetByID(id)
}

func (s *stockOpnameService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *stockOpnameService) IsExist(id uint) (bool, error) {
	return s.repo.IsExist(id)
}
