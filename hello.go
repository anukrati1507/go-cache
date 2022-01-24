package main

import (
	"time"
)

type Data struct {
	data    int
	expTime int64
}
type InMap struct {
	Map map[string]Data
}
type Parameters struct {
	data int
	key  string
	ttl  int64
}

func (c *InMap) set(param Parameters) {
	if param.ttl == 0 {
		param.ttl = 60
	} else if param.ttl < 0 {
		panic("Time cannot be a negative entity!")
	}
	c.Map[param.key] = Data{data: param.data, expTime: time.Now().Unix() + param.ttl}
}
func (c *InMap) get(key string) (value int) {
	var state bool
	var data1 Data
	data1, state = c.Map[key]
	value = data1.data
	if !state {
		panic("Key is not found in the cache.")
	}
	return
}

/*func get(key int) {
	if(m[key]) {
		fmt.println(m[key])
	}
	else {
		fmt.println("Key not found!")
	}
}
func set(key int,value int,time int) {
	if(time==nil){
		time=60
	}
}*/
func (c *InMap) delete(key string) {
	var state bool
	_, state = c.Map[key]
	if !state {
		panic("Error, Key not found in the map.")
	}
	delete(c.Map, key)
}
func newCache() (inmap InMap) {
	inmap = InMap{Map: make(map[string]Data)}
	return
}
func checkExpiry(inmap *InMap) {
	for {
		for key, value := range inmap.Map {
			if value.expTime < time.Now().Unix() {
				delete(inmap.Map, key)
			}
		}
	}
}
