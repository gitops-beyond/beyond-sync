package models

import (
	"github.com/gitops-beyond/beyond-sync/internal/redis"
)

type SyncResponse struct {
    Timestamp string     `json:"timestamp"`
    Data      redis.RedisValue `json:"data"`
}

type SyncListResponse []SyncResponse