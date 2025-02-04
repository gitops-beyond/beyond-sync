package routes

import( 
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/handlers"
)

func LoadRoutes(r *gin.Engine){
	r.GET("/sync", handlers.GetAllSyncs)
	r.GET("/sync/:timestamp", handlers.GetSyncByDate)
}