package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/routes"
)

func StartServer() {
	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run(":8080")
}
