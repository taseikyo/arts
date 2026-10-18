/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-09-10 20:04:24
 * @link    github.com/taseikyo
 */

package locker

import (
	"context"
)

var lockers map[string]map[uint64]chan interface{}
var addLock chan lockImpl
var delLock chan lockImpl
var tranceId uint64

type lockImpl struct {
	delId uint64
	addId chan uint64
	wait  chan interface{}
	key   string
}

func init() {
	lockers = make(map[string]map[uint64]chan interface{})
	addLock = make(chan lockImpl)
	delLock = make(chan lockImpl)
	run()
}

func run() {
	go func() {
		var (
			ok1, ok2     bool
			addLs, delLs map[uint64]chan interface{}
		)

		for {
			select {
			case add := <-addLock:
				if addLs, ok1 = lockers[add.key]; !ok1 {
					addLs = make(map[uint64]chan interface{})
				}
				tranceId++
				addLs[tranceId] = add.wait
				lockers[add.key] = addLs
				add.addId <- tranceId
				close(add.addId)
				if len(addLs) == 1 {
					add.wait <- struct{}{}
					close(add.wait)
				}

			case del := <-delLock:
				if delLs, ok2 = lockers[del.key]; ok2 {
					delete(delLs, del.delId)
					for _, v := range delLs {
						v <- struct{}{}
						close(v)
						break
					}
				}

			}
		}

	}()
}

type Memory struct {
	key string
	id  uint64
}

func (m *Memory) SLock(ctx context.Context) (timeOut bool, err error) {

	l := lockImpl{
		addId: make(chan uint64, 1),
		wait:  make(chan interface{}, 1),
		key:   m.key,
	}
	addLock <- l
	m.id = <-l.addId

	select {
	case <-l.wait:
		return
	case <-ctx.Done():
		timeOut = true
		return
	}

}

func (m *Memory) UnSLock() (err error) {
	l := lockImpl{
		delId: m.id,
		key:   m.key,
	}
	delLock <- l

	return
}

func NewMemory(path string) *Memory {
	return &Memory{
		key: path,
	}
}
