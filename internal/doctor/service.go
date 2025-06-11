package doctor

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrNotFound     = errors.New("dokter tidak ditemukan")
	ErrInvalidInput = errors.New("input tidak valid atau tidak lengkap")
	ErrSTRExists    = errors.New("nomor STR sudah digunakan oleh dokter lain yang aktif")
)

type Service interface {
	CreateDoctor(doctor *Doctor) (*Doctor, error)
	GetAllDoctors(searchQuery string) ([]Doctor, error)
	GetDoctorByID(id uint) (*Doctor, error)
	UpdateDoctor(id uint, doctor *Doctor) (*Doctor, error)
	DeleteDoctor(id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) CreateDoctor(doctor *Doctor) (*Doctor, error) {
	doctor.FullName = strings.TrimSpace(doctor.FullName)
	if doctor.FullName == "" || doctor.Specialization == "" || doctor.PhoneNumber == "" {
		return nil, ErrInvalidInput
	}

	if doctor.STRNumber != "" {
		existing, err := s.repository.FindActiveBySTR(doctor.STRNumber)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing != nil {
			return nil, ErrSTRExists
		}
	}

	if doctor.Status == "" {
		doctor.Status = "Aktif"
	}

	return s.repository.Create(doctor)
}

func (s *service) UpdateDoctor(id uint, doctor *Doctor) (*Doctor, error) {
	if _, err := s.repository.GetByID(id); err != nil {
		return nil, err
	}

	if doctor.STRNumber != "" {
		existing, err := s.repository.FindActiveBySTR(doctor.STRNumber)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing != nil && existing.ID != id {
			return nil, ErrSTRExists
		}
	}

	return s.repository.Update(id, doctor)
}

func (s *service) GetAllDoctors(searchQuery string) ([]Doctor, error) {
	return s.repository.GetAll(searchQuery)
}

func (s *service) GetDoctorByID(id uint) (*Doctor, error) {
	return s.repository.GetByID(id)
}

func (s *service) DeleteDoctor(id uint) error {
	if _, err := s.repository.GetByID(id); err != nil {
		return err
	}
	return s.repository.Delete(id)
}
