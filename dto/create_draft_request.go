package dto

import "time"

// Request and response structures

type CreateDraftRequest struct {
	OpnameDate time.Time `json:"opname_date" binding:"required"`
	Notes      string    `json:"notes"`
}

type UpdateDraftRequest struct {
	OpnameDate time.Time `json:"opname_date" binding:"required"`
	Notes      string    `json:"notes"`
}
