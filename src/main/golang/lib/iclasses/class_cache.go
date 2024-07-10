package iclasses

import (
	"sync"

	"github.com/starter-go/keyvalues"
)

// ClassCache ...
type ClassCache struct {

	//starter:component

	_as func(keyvalues.ClassGetter) //starter:as("#")

	Loader keyvalues.ClassLoader //starter:inject("#")

	cache *innerClassCache
}

func (inst *ClassCache) _impl() keyvalues.ClassGetter {
	return inst
}

func (inst *ClassCache) getCache() *innerClassCache {
	c := inst.cache
	if c == nil {
		c = &innerClassCache{}
		c.init()
		inst.cache = c
	}
	return c
}

// GetClass ...
func (inst *ClassCache) GetClass(name keyvalues.ClassName) (keyvalues.Class, error) {
	c := inst.getCache()
	item := c.get(name)
	if item == nil {
		item2, err := inst.Loader.LoadClass(name)
		if err != nil {
			return nil, err
		}
		item = item2
		c.put(name, item)
	}
	return item, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerClassCache struct {
	mutex sync.Mutex
	table map[keyvalues.ClassName]keyvalues.Class
}

func (inst *innerClassCache) init() {
	inst.table = make(map[keyvalues.ClassName]keyvalues.Class)
}

func (inst *innerClassCache) put(name keyvalues.ClassName, item keyvalues.Class) {

	locker := &inst.mutex
	locker.Lock()
	locker.Unlock()

	inst.table[name] = item
}

func (inst *innerClassCache) get(name keyvalues.ClassName) keyvalues.Class {

	locker := &inst.mutex
	locker.Lock()
	locker.Unlock()

	return inst.table[name]
}
