package main

import (
	"fmt"
	"time"
)

type Data struct {
	data           int64
	expirationTime int64
	index          int
}

type InMap struct {
	Map map[string]*Data
}

type Parameters struct {
	data int64
	key  string
	ttl  int64
}

func (c *InMap) set(param Parameters) (item *Data, exists bool) {
	if param.ttl == 0 {
		param.ttl = 60
	} else if param.ttl < 0 {
		panic("Time cannot be a negative entity")
	}
	_, exists = c.Map[param.key]
	c.Map[param.key] = &Data{data: param.data, expirationTime: time.Now().Unix() + param.ttl}
	item = c.Map[param.key]
	return
}

func (c *InMap) get(key string) (value int64) {
	var state bool
	var item *Data
	item, state = c.Map[key]
	value = item.data
	if !state {
		panic("Key not found")
	}
	return
}

func (c *InMap) delete(key string) (item *Data) {
	var state bool
	item, state = c.Map[key]
	if !state {
		panic("Key not found")
	}
	delete(c.Map, key)
	return
}

func (c *InMap) printMap() {
	for key, value := range c.Map {
		fmt.Printf("%s: %d %d \n", key, value.data, value.expirationTime)
	}
}

/*func (c *InMap) startExpirationTime() {
	timer := time.NewTimer(time.Hour)
	for {
		var sleepTime time.Duration
		//cache.mutex.Lock()
		//cache.hasNotified = false
		if c.priorityQueue.Len() > 0 {
			sleepTime = time.Until(cache.priorityQueue.root().expireAt)
			if sleepTime < 0 && cache.priorityQueue.root().expireAt.IsZero() {
				sleepTime = time.Hour
			} else if sleepTime < 0 {
				sleepTime = time.Microsecond
			} else {
			sleepTime = time.Hour
		}

		//cache.expirationTime = time.Now().Add(sleepTime)
		//cache.mutex.Unlock()

		timer.Reset(sleepTime)
		select {
		case shutdownFeedback := <-cache.shutdownSignal:
			timer.Stop()
			//cache.mutex.Lock()
			if cache.priorityQueue.Len() > 0 {
				cache.evictjob(Closed)
			}
			cache.mutex.Unlock()
			shutdownFeedback <- struct{}{}
			return
		case <-timer.C:
			timer.Stop()
			cache.mutex.Lock()
			if cache.priorityQueue.Len() == 0 {
				cache.mutex.Unlock()
				continue
			}

			cache.cleanjob()
			cache.mutex.Unlock()

		case <-cache.expirationNotification:
			timer.Stop()
			continue
		}
	}
}
*/

func newCache() (inmap InMap) {
	inmap = InMap{Map: make(map[string]*Data)}
	return
}
