/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-09-10 20:03:23
 * @link    github.com/taseikyo
 */

package locker

import (
	"context"
	"github.com/go-redis/redis"
	"runtime"
	"time"
)

type Lock struct {
	resource string
	value    interface{}
	timeout  time.Duration
	redisCli *redis.ClusterClient //这个是链接redis集群的cli，可以自行修改
}

func NewRedis(redisCli *redis.ClusterClient, resource string, value interface{}, timeOut time.Duration) *Lock {
	return &Lock{
		resource: resource,
		value:    value,
		timeout:  timeOut,
		redisCli: redisCli,
	}
}

func (lock *Lock) TryLock() (ok bool, err error) {
	ok, err = lock.redisCli.SetNX(lock.resource, lock.value, lock.timeout).Result()
	//log.Printf("resource:%s, timeout:%v, ok:%v, err:%v\n", lock.resource, lock.timeout, ok, err)

	return
}

func (lock *Lock) Unlock() (err error) {
	err = lock.redisCli.Del(lock.resource).Err()
	return
}

func (lock *Lock) SpinLockUntilTimeOut(ctx context.Context, d time.Duration) (timeOut bool, err error) {
	var (
		now time.Time
		ok  bool
	)

	endTime := time.Now().Add(d)
	for {
		select {
		case <-ctx.Done():
			timeOut = true
			//log.Printf("SpinLockUntilTimeOut at ctx.Done()")
			return

		default:
			now = time.Now()
			if now.After(endTime) {
				timeOut = true
				//log.Printf("SpinLockUntilTimeOut at d")
				return
			}

			ok, err = lock.TryLock()
			if err != nil {
				return
			}

			if ok {
				return
			} else {
				runtime.Gosched()
			}

		}

	}

}
