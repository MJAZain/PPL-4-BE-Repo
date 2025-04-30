package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`                         // contoh: "admin", "user"
	NIP      string `gorm:"column:nip;unique" json:"nip"` // Nomor Induk Pegawai, diasumsikan unik
}
