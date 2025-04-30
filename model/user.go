package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `gorm:"unique" json:"phone"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`                         // contoh: "admin", "user"
	NIP       string    `gorm:"column:nip;unique" json:"nip"` // Nomor Induk Pegawai, diasumsikan unik
	Active    bool      `json:"active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
