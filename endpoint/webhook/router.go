package webhook

import (
	"github.com/farwydi/cleanwhale/tonic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"telegram-state/domain"
)

func NewRouter(cfg domain.Config) http.Handler {
	r := tonic.NewMix(cfg.Project.Mode, zap.L().Named("webhook"))

	v1 := r.Group("/v1", tonic.V("v1"))
	{
		v1.GET("/webhook", func(c *gin.Context) {
			c.Set("handler.url", "/webhook")
		})
	}

	return r
}
