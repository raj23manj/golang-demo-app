package dao

import (
	"fmt"

	"github.com/raj23manj/demo-app-golang/db/connection"
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/dto"
)

type tenantDao struct{}

type tenantDaoInterface interface {
	Create(tenant *dto.CreateTenantRequest) (*domain.Tenant, *dto.DtoErrorResponse)
	GetTenant(id uint64) (*domain.Tenant, *dto.DtoErrorResponse)
}

var (
	TenantDao tenantDaoInterface
)

func init() {
	TenantDao = &tenantDao{}
}

func (t *tenantDao) Create(tenant *dto.CreateTenantRequest) (*domain.Tenant, *dto.DtoErrorResponse) {
	newTenant := &domain.Tenant{
		Name:      tenant.Name,
		Active:    tenant.Active,
		ExpiresIn: tenant.ExpiresIn,
	}
	result := connection.DB.Create(newTenant)
	if result.Error != nil {
		fmt.Println(result.Error)

		return nil, &dto.DtoErrorResponse{
			Message: "error creating records !!!!",
		}
	}

	return newTenant, nil
}

func (t *tenantDao) GetTenant(id uint64) (*domain.Tenant, *dto.DtoErrorResponse) {
	tenant := domain.Tenant{}
	result := connection.DB.Where("id = ?", id).First(&tenant)
	if result.Error != nil {
		return nil, &dto.DtoErrorResponse{
			Message: fmt.Sprintf("unable to find tenant %v", id),
		}
	}

	return &tenant, nil
}
