// /internal/models/message.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Content   string
	CreatedAt time.Time
}
