package middleware

import (
	"github.com/Thanga-tamil/noway_service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		tenantId := c.Request.Header.Get("tenant-x")

		db := config.TenantDBs[tenantId]
		logrus.Info("Pad tenant db for domain: ", tenantId)

		c.Set(tenantId, db)
		c.Next()


		// after request
		// todo


	}
}

