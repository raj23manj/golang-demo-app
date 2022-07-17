package dto

import "time"

type CreateTenantRequest struct {
	Name      string                 `json:"name"`
	Active    bool                   `json:"active"`
	ExpiresIn map[string]interface{} `json:"expires_in"`
}

type TenantMapperDto struct {
	Id        uint64                 `json:"id"`
	Name      string                 `json:"name"`
	Active    bool                   `json:"active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	ExpiresIn map[string]interface{} `json:"expires_in"`
}
