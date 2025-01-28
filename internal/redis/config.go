package redis

import (

)

type RedisValue struct {
    Sha string `json:"sha"`
    Status string `json:"status"`
    Message string `json:"message"`
}