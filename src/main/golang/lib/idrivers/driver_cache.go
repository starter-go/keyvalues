package idrivers

import (
	"sync"

	"github.com/starter-go/keyvalues"
)

// DriverCache ...
type DriverCache struct {

	//starter:component

	_as func(keyvalues.DriverGetter) //starter:as("#")

	Loader DriverLoader //starter:inject("#")

	cache *innerDriverCache
}

func (inst *DriverCache) _impl() keyvalues.DriverGetter {
	return inst
}

func (inst *DriverCache) getCache() *innerDriverCache {
	c := inst.cache
	if c == nil {
		c = new(innerDriverCache)
		c.init()
		inst.cache = c
	}
	return c
}

// GetDriver ...
func (inst *DriverCache) GetDriver(name string) (keyvalues.Driver, error) {
	c := inst.getCache()
	item := c.get(name)
	if item == nil {
		item2, err := inst.Loader.Load(name)
		if err != nil {
			return nil, err
		}
		item = item2
		c.put(name, item)
	}
	return item, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerDriverCache struct {
	mutex sync.Mutex
	table map[string]keyvalues.Driver
}

func (inst *innerDriverCache) init() {
	inst.table = make(map[string]keyvalues.Driver)
}

func (inst *innerDriverCache) put(name string, item keyvalues.Driver) {
	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()
	inst.table[name] = item
}

func (inst *innerDriverCache) get(name string) keyvalues.Driver {
	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()
	return inst.table[name]
}
