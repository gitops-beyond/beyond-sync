package routes

import( 
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/controllers"
)

func LoadRoutes(r *gin.Engine){
	r.GET("/sync", controllers.GetSyncs)
}