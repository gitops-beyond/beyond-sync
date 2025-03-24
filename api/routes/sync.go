package routes

import( 
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/api/handlers"
)

// LoadRoutes configures the API endpoints for sync operations
func LoadRoutes(r *gin.Engine){
	r.GET("/sync", handlers.GetAllSyncs)         // Get all sync records
	r.GET("/sync/:timestamp", handlers.GetSyncByDate) // Get sync by specific timestamp
	r.POST("/sync/trigger", handlers.TriggerSync)     // Trigger new sync operation
}