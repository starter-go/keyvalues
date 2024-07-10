package ram

import (
	"fmt"
	"sync"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/vlog"
)

// MemoryDriver ...
type MemoryDriver struct {

	//starter:component

	_as func(keyvalues.Registry) //starter:as(".")
}

func (inst *MemoryDriver) _impl() (keyvalues.Registry, keyvalues.Driver) {
	return inst, inst
}

// ListRegistrations ...
func (inst *MemoryDriver) ListRegistrations() []*keyvalues.Registration {

	r1 := new(keyvalues.Registration)
	info := new(keyvalues.DriverInfo)

	r1.Driver = info
	r1.Enabled = true
	r1.Priority = 0

	info.Driver = inst
	info.Name = "memory"

	return []*keyvalues.Registration{r1}
}

// Open ...
func (inst *MemoryDriver) Open(cfg *keyvalues.Configuration) (keyvalues.Store, error) {
	st := new(memoryStore)
	err := st.open(cfg)
	if err != nil {
		return nil, err
	}
	return st, nil
}

////////////////////////////////////////////////////////////////////////////////

type memoryStoreItem struct {
	key   string
	value []byte
	t1    lang.Time
	t2    lang.Time
}

////////////////////////////////////////////////////////////////////////////////

type memoryStore struct {
	table  map[string]*memoryStoreItem
	mutex  sync.Mutex
	config *keyvalues.Configuration
}

func (inst *memoryStore) _impl() keyvalues.Store {
	return inst
}

func (inst *memoryStore) open(cfg *keyvalues.Configuration) error {
	inst.table = make(map[string]*memoryStoreItem)
	inst.config = cfg
	return nil
}

func (inst *memoryStore) Put(key string, value []byte, opt *keyvalues.Options) error {

	if value == nil {
		value = make([]byte, 0)
	}
	if opt == nil {
		opt = new(keyvalues.Options)
	}
	if opt.MaxAge < 1 {
		opt.MaxAge = inst.config.DefaultMaxAge
	}

	now := lang.Now()
	item := new(memoryStoreItem)
	item.key = key
	item.t1 = now
	item.t2 = now.Add(opt.MaxAge)
	item.value = value

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	inst.table[key] = item
	return nil
}

func (inst *memoryStore) Get(key string) ([]byte, error) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	item := inst.table[key]
	ok := inst.isItemOK(item)
	if !ok {
		msg := "key-value not found"
		vlog.Warn(msg+", key=%s", key)
		return nil, fmt.Errorf(msg)
	}
	return item.value, nil
}

func (inst *memoryStore) Contains(key string) (bool, error) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	item := inst.table[key]
	ok := inst.isItemOK(item)
	return ok, nil
}

func (inst *memoryStore) isItemOK(item *memoryStoreItem) bool {

	if item == nil {
		return false
	}

	if item.value == nil {
		return false
	}

	now := lang.Now()
	return (item.t1 <= now) && (now <= item.t2)
}

////////////////////////////////////////////////////////////////////////////////
