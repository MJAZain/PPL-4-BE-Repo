package patient

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound       = errors.New("pasien tidak ditemukan")
	ErrInvalidInput   = errors.New("input tidak valid atau tidak lengkap")
	ErrIdentityExists = errors.New("nomor identitas sudah digunakan oleh pasien lain yang aktif")
)

type Service interface {
	CreatePatient(patient *Patient) (*Patient, error)
	GetAllPatients(searchQuery string) ([]Patient, error)
	GetPatientByID(id uint) (*Patient, error)
	UpdatePatient(id uint, patient *Patient) (*Patient, error)
	DeletePatient(id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func calculateAge(birthdate time.Time) int {
	now := time.Now()
	age := now.Year() - birthdate.Year()
	if now.YearDay() < birthdate.YearDay() {
		age--
	}
	return age
}

func (s *service) enrichPatientData(patient *Patient) {
	if patient != nil {
		patient.Age = calculateAge(patient.DateOfBirth)
	}
}

func (s *service) CreatePatient(patient *Patient) (*Patient, error) {
	patient.FullName = strings.TrimSpace(patient.FullName)
	if patient.FullName == "" || patient.Gender == "" || patient.PlaceOfBirth == "" || patient.DateOfBirth.IsZero() {
		return nil, ErrInvalidInput
	}

	if patient.IdentityNumber != "" {
		existing, err := s.repository.FindActiveByIdentityNumber(patient.IdentityNumber)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing != nil {
			return nil, ErrIdentityExists
		}
	}

	if patient.Status == "" {
		patient.Status = "Aktif"
	}

	newPatient, err := s.repository.Create(patient)
	if err != nil {
		return nil, err
	}
	s.enrichPatientData(newPatient)
	return newPatient, nil
}

func (s *service) UpdatePatient(id uint, patient *Patient) (*Patient, error) {
	if _, err := s.repository.GetByID(id); err != nil {
		return nil, err
	}

	if patient.IdentityNumber != "" {
		existing, err := s.repository.FindActiveByIdentityNumber(patient.IdentityNumber)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing != nil && existing.ID != id {
			return nil, ErrIdentityExists
		}
	}

	updatedPatient, err := s.repository.Update(id, patient)
	if err != nil {
		return nil, err
	}
	s.enrichPatientData(updatedPatient)
	return updatedPatient, nil
}

func (s *service) GetAllPatients(searchQuery string) ([]Patient, error) {
	patients, err := s.repository.GetAll(searchQuery)
	if err != nil {
		return nil, err
	}
	for i := range patients {
		s.enrichPatientData(&patients[i])
	}
	return patients, nil
}

func (s *service) GetPatientByID(id uint) (*Patient, error) {
	patient, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	s.enrichPatientData(patient)
	return patient, nil
}

func (s *service) DeletePatient(id uint) error {
	if _, err := s.repository.GetByID(id); err != nil {
		return err
	}
	return s.repository.Delete(id)
}
