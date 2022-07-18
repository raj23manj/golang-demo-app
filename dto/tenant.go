package dto

import "time"

type CreateTenantRequest struct {
	Name      string                 `json:"name" binding:"required"`
	Active    bool                   `json:"active" binding:"required"`
	ExpiresIn map[string]interface{} `json:"expires_in" binding:"required"`
}

type TenantMapperDto struct {
	Id        uint64                 `json:"id"`
	Name      string                 `json:"name"`
	Active    bool                   `json:"active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	ExpiresIn map[string]interface{} `json:"expires_in"`
}

type TenantGetRequest struct {
	TenantId uint64 `uri:"tenant_id" binding:"required"`
}
