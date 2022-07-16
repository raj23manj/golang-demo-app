package services

import (
	"net/http"

	"github.com/raj23manj/demo-app-golang/db/connection"
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

type tenantService struct{}

type tenantServiceInterface interface {
	GetTenants() (*[]domain.Tenant, errors.ApiError)
}

var (
	TenantService tenantServiceInterface
)

func init() {
	TenantService = &tenantService{}
}

func (t *tenantService) GetTenants() (*[]domain.Tenant, errors.ApiError) {
	tenants := []domain.Tenant{}
	result := connection.DB.Find(&tenants)
	if result.Error != nil {
		return nil, errors.NewApiError(http.StatusNotFound, "not able to retrive all records!!!")
	}

	return &tenants, nil
}
