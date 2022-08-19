package healthz

import (
	"github.com/STommydx/togolist/internal/pkg/router"
	"github.com/gin-gonic/gin"
)

type HealthzController struct {
	healthzService *HealthzService
}

func NewHealthzController(healthzService *HealthzService) *HealthzController {
	return &HealthzController{healthzService}
}

func registerHealthzRoutes(hc *HealthzController, r *router.Router) {
	r.RegisterApiRoutes("/healthz", func(rg *gin.RouterGroup) {
		rg.GET("", hc.healthCheck)
	})
}

// healthCheck godoc
// @Summary      Health Checking
// @Description  Health Checking for API services
// @Produce      json
// @Success      200  {object}  dto.HealthzResult
// @Router       /healthz [get]
func (hc *HealthzController) healthCheck(c *gin.Context) {
	result, err := hc.healthzService.HealthCheck()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	if result.Status == "ok" {
		c.JSON(200, result)
	} else {
		c.JSON(503, result)
	}
}
