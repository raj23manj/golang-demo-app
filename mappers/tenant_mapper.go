package mappers

import (
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/dto"
)

func MapTenantData(tenant *domain.Tenant) *dto.TenantMapperDto {
	resp := &dto.TenantMapperDto{
		Id:        uint64(tenant.ID),
		Name:      tenant.Name,
		Active:    tenant.Active,
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
		ExpiresIn: tenant.ExpiresIn,
	}

	return resp
}
