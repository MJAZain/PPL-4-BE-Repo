// model/user_log.go
package model

import "time"

type UserLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Activity  string    `json:"activity"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"` // optional: join ke user
}
