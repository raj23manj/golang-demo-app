package mappers

import (
	"time"

	"github.com/raj23manj/demo-app-golang/domain"
)

type tenantMapper struct {
	Id        uint64                 `json:"id"`
	Name      string                 `json:"name"`
	Active    bool                   `json:"active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	ExpiresIn map[string]interface{} `json:"expires_in"`
}

type tenantMapperInterface interface {
	MapData(tenant *domain.Tenant) *tenantMapper
}

var (
	TenantMapper tenantMapperInterface
)

func init() {
	TenantMapper = &tenantMapper{}
}

func (t *tenantMapper) MapData(tenant *domain.Tenant) *tenantMapper {
	resp := &tenantMapper{
		Id:        uint64(tenant.ID),
		Name:      tenant.Name,
		Active:    tenant.Active,
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
		ExpiresIn: tenant.ExpiresIn,
	}

	return resp
}
