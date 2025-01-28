package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

func GetSyncs(c *gin.Context) {
	redis.GetAllSyncRecords()
}