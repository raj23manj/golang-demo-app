package domain

import (
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name      string
	Active    bool
	ExpiresIn JSONMap `gorm:"type:jsonb;default:'null';"`
}
