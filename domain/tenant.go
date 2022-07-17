package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name      string
	Active    bool
	ExpiresIn JSONMap `gorm:"type:jsonb;default:'null';"`
}

type JSONMap map[string]interface{}

func (t JSONMap) Value() (driver.Value, error) {
	fmt.Println("data value %v", t)
	return json.Marshal(t)
}

func (t *JSONMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	res := json.Unmarshal(b, &t)
	fmt.Println("data scan %v", res)
	return json.Unmarshal(b, &t)
}
