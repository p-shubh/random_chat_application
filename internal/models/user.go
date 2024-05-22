// /internal/models/user.go
package models

import "github.com/google/uuid"

type User struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
    Username string    `gorm:"type:varchar(100)"`
}

