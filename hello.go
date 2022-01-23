package main

import (
	"time"
	"fmt"
)

type Data struct {
	data int
	expTime int
}

type InMap struct {
	Map map[string]Value
}

type Param struct {
	data int,
	key string
	ttl int
}

func (c *InMap) set(param Param) {
	if param.ttl == nil {
		param.ttl=60
	} else if param.ttl<0 {
		panic("Time cannot be a negative entity!")
	}
	c.Map[param.key] = Data{data: param.data, expTime: time.Now().Unix()+param.ttl}
}

func (c *InMap) get(key String) {
	var state bool
	var data1 Data
	data1, state = c.Map[key]
	value = item.data
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
func (c *InMap) delete(key String) {
	var state bool
	_, state = c.Map[key]
	if !state {
		panic("ERR: Key not found.")
	}
	delete(c.Map, key)
}

func newCache() (inmap InMap) {
	inmap = InMap{Map: make(map[string]Data)}
	return
}

func checkExpiry(inmap *InMap){
	for {
		for key, value := range inmap.Map {
			if value.expirationTime < time.Now().Unix() {
				delete(inmap.Map,key)
			}
		}
	}
}
