package dao_redis

import (
	"context"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/google/uuid"
	"log"
	"time"
)

// AcquireLock 获取分布式锁
// 参数:
//   - key: 锁的键名
//   - acquireTimeout: 获取锁的超时时间
//   - lockTimeout: 锁的自动过期时间
//
// 返回值:
//   - string: 锁的唯一标识(成功时),空字符串(失败时)
//   - bool: 是否获取成功
func AcquireLock(key string, acquireTimeout, lockTimeout time.Duration) (string, bool) {
	if acquireTimeout <= 0 {
		return "", false
	}
	endTime := time.Now().Add(acquireTimeout)
	//锁值
	lockValue := uuid.NewString()

	for time.Now().Before(endTime) {
		// 使用SETNX原子操作尝试获取锁
		set, err := global.RDB.SetNX(context.Background(), key, lockValue, lockTimeout).Result()
		if err != nil {
			return "", false
		}

		if set {
			return lockValue, true
		}
		// 短暂休眠避免频繁重试
		time.Sleep(100 * time.Millisecond)
	}

	return "", false
}

// ReleaseLock 释放分布式锁
// 参数:
//   - key: 锁的键名
//   - lockValue: 锁的唯一标识
//
// 返回值:
//   - bool: 是否释放成功
func ReleaseLock(key, lockValue string) bool {
	if key == "" || lockValue == "" {
		return false
	}

	// 使用Lua脚本确保原子性

	luaScript := `
    if redis.call("GET",KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end
`
	result, err := global.RDB.Eval(context.Background(), luaScript, []string{key}, lockValue).Result()
	if err != nil {
		log.Printf("释放锁时出错: %v", err)
		return false
	}
	return result.(int64) == 1
}
