package domain

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name      string
	Active    bool
	ExpiresIn pgtype.JSONB `gorm:"type:jsonb;default:'null';"`
}
