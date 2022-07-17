package tenant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/dto"
	"github.com/raj23manj/demo-app-golang/services"
	"github.com/raj23manj/demo-app-golang/utils/controller"
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

	controller.Respond(c, http.StatusOK, newTenant)
}
