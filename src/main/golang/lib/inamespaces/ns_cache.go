package inamespaces

import (
	"sync"

	"github.com/starter-go/keyvalues"
)

// NamespaceCache ...
type NamespaceCache struct {

	//starter:component

	_as func(keyvalues.NamespaceGetter) //starter:as("#")

	Loader keyvalues.NamespaceLoader //starter:inject("#")

	cache *innerNamespaceCache
}

func (inst *NamespaceCache) _impl() keyvalues.NamespaceGetter {
	return inst
}

func (inst *NamespaceCache) getCache() *innerNamespaceCache {
	c := inst.cache
	if c == nil {
		c = new(innerNamespaceCache)
		c.init()
		inst.cache = c
	}
	return c
}

// GetNamespace ...
func (inst *NamespaceCache) GetNamespace(ns keyvalues.NS) (keyvalues.Namespace, error) {
	c := inst.getCache()
	item := c.get(ns)
	if item == nil {
		item2, err := inst.Loader.LoadNamespace(ns)
		if err != nil {
			return nil, err
		}
		item = item2
		c.put(ns, item)
	}
	return item, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerNamespaceCache struct {
	mutex sync.Mutex
	table map[keyvalues.NS]keyvalues.Namespace
}

func (inst *innerNamespaceCache) init() {
	inst.table = make(map[keyvalues.NS]keyvalues.Namespace)
}

func (inst *innerNamespaceCache) get(name keyvalues.NS) keyvalues.Namespace {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	return inst.table[name]
}

func (inst *innerNamespaceCache) put(name keyvalues.NS, item keyvalues.Namespace) {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	inst.table[name] = item
}
