package istores

import (
	"sync"

	"github.com/starter-go/keyvalues"
)

// StoreCache ...
type StoreCache struct {

	//starter:component

	_as func(keyvalues.StoreGetter) //starter:as("#")

	Loader keyvalues.StoreLoader //starter:inject("#")

	cache *innerStoreCache
}

func (inst *StoreCache) _impl() keyvalues.StoreGetter {
	return inst
}

func (inst *StoreCache) getCache() *innerStoreCache {
	c := inst.cache
	if c == nil {
		c = new(innerStoreCache)
		c.init()
		inst.cache = c
	}
	return c
}

// GetStore ...
func (inst *StoreCache) GetStore(name string) (keyvalues.Store, error) {
	c := inst.getCache()
	item1 := c.get(name)
	if item1 == nil {
		item2, err := inst.Loader.LoadStore(name)
		if err != nil {
			return nil, err
		}
		item1 = item2
		c.put(name, item2)
	}
	return item1, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerStoreCache struct {
	table map[string]keyvalues.Store
	mutex sync.Mutex
}

func (inst *innerStoreCache) init() {
	inst.table = make(map[string]keyvalues.Store)
}

func (inst *innerStoreCache) put(name string, item keyvalues.Store) {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	inst.table[name] = item
}

func (inst *innerStoreCache) get(name string) keyvalues.Store {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	return inst.table[name]
}
