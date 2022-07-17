package tenant

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/dto"
	"github.com/raj23manj/demo-app-golang/mappers"
	"github.com/raj23manj/demo-app-golang/services"
	"github.com/raj23manj/demo-app-golang/utils/controller"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

func Create(c *gin.Context) {
	tenant := &dto.CreateTenantRequest{
		Name:   "new_tenant",
		Active: true,
		ExpiresIn: map[string]interface{}{
			"days": 5,
		},
	}

	newTenant, apiErr := services.TenantService.Create(tenant)
	if apiErr != nil {
		controller.RespondError(c, apiErr)
		return
	}

	mappedResult := mappers.MapTenantData(newTenant)
	controller.Respond(c, http.StatusOK, mappedResult)
}

func GetTenant(c *gin.Context) {
	tenantIdParam := c.Param("tenant_id")
	tenantId, err := (strconv.ParseInt(tenantIdParam, 10, 64))
	if err != nil {
		apiErr := errors.NewBadRequestError("tenant id must be a number!!!")
		controller.RespondError(c, apiErr)
		return
	}

	tenant, apiErr := services.TenantService.GetTenant(uint64(tenantId))
	if apiErr != nil {
		controller.RespondError(c, apiErr)
		return
	}
	mappedResult := mappers.MapTenantData(tenant)
	controller.Respond(c, http.StatusOK, mappedResult)
}

func GetTenants(c *gin.Context) {
	tenants, apiErr := services.TenantService.GetTenants()

	if apiErr != nil {
		controller.RespondError(c, apiErr)
		return
	}

	fmt.Println("tenants %v", reflect.TypeOf(tenants))

	fmt.Println("tenants %v", *tenants)

	// mappedTenants := funk.Map(*tenants, func(tenant *domain.Tenant) *dto.TenantMapperDto {
	// 	return mappers.MapTenantData(tenant)
	// })

	//mappedTenants := app.MapTo(tenants, func())

	mapped := make([]*dto.TenantMapperDto, len(*tenants))

	for i, t := range *tenants {
		mapped[i] = mappers.MapTenantData(&t)
	}

	controller.Respond(c, http.StatusOK, mapped)
}
