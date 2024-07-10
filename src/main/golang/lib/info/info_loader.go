package info

import (
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/keyvalues"
)

// GetterImpl ...
type GetterImpl struct {

	//starter:component

	_as func(Getter) //starter:as("#")

	AppContext application.Context  //starter:inject("context")
	RegList    []keyvalues.Registry //starter:inject(".")

	cache *cache
	mutex sync.Mutex
}

func (inst *GetterImpl) _impl() Getter {
	return inst
}

// GetClassInfoList ...
func (inst *GetterImpl) GetClassInfoList() []*keyvalues.ClassInfo {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	c := inst.getCache()
	src := c.classes
	dst := make([]*keyvalues.ClassInfo, 0)

	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

// GetNamespaceInfoList ...
func (inst *GetterImpl) GetNamespaceInfoList() []*keyvalues.NamespaceInfo {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	c := inst.getCache()
	src := c.namespaces
	dst := make([]*keyvalues.NamespaceInfo, 0)

	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

// GetStoreInfoList ...
func (inst *GetterImpl) GetStoreInfoList() []*keyvalues.StoreInfo {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	c := inst.getCache()
	src := c.stores
	dst := make([]*keyvalues.StoreInfo, 0)

	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

// GetDriverInfoList ...
func (inst *GetterImpl) GetDriverInfoList() []*keyvalues.DriverInfo {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	c := inst.getCache()
	src := c.drivers
	dst := make([]*keyvalues.DriverInfo, 0)

	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

func (inst *GetterImpl) getCache() *cache {
	c := inst.cache
	if c == nil {
		c = inst.loadCache()
		inst.cache = c
	}
	return c
}

func (inst *GetterImpl) loadCache() *cache {
	ldr := &loader{parent: inst}
	c, err := ldr.load()
	if err != nil {
		panic(err)
	}
	return c
}

////////////////////////////////////////////////////////////////////////////////

type cache struct {
	classes    map[keyvalues.ClassName]*keyvalues.ClassInfo
	namespaces map[keyvalues.NS]*keyvalues.NamespaceInfo
	stores     map[string]*keyvalues.StoreInfo
	drivers    map[string]*keyvalues.DriverInfo
}

////////////////////////////////////////////////////////////////////////////////

type loader struct {
	parent *GetterImpl
}

func (inst *loader) load() (*cache, error) {

	c := &cache{}
	c.classes = make(map[keyvalues.ClassName]*keyvalues.ClassInfo)
	c.drivers = make(map[string]*keyvalues.DriverInfo)
	c.namespaces = make(map[keyvalues.NS]*keyvalues.NamespaceInfo)
	c.stores = make(map[string]*keyvalues.StoreInfo)

	err := inst.loadClassesAndNS(c)
	if err != nil {
		return nil, err
	}

	err = inst.loadStores(c)
	if err != nil {
		return nil, err
	}

	err = inst.loadDrivers(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (inst *loader) loadClassesAndNS(c *cache) error {
	const (
		prefix = "key-value-class."
		suffix = ".name"
	)
	ac := inst.parent.AppContext
	g := ac.GetProperties().Getter().Required()

	ids := g.ListItems(prefix, suffix)
	for _, id := range ids {
		keyx := prefix + id + "."

		enabled := g.GetBool(keyx + "enabled")
		maxage := g.GetInt64(keyx + "default-max-age")
		namespace := g.GetString(keyx + "namespace")
		store := g.GetString(keyx + "store")
		name := g.GetString(keyx + "name")

		alias := keyvalues.ClassAlias(name)
		ns := keyvalues.NS(namespace)
		fullname := ns.GetClassName(alias)

		iClass := &keyvalues.ClassInfo{}
		iNS := &keyvalues.NamespaceInfo{}

		iClass.ID = id
		iClass.Enabled = enabled
		iClass.MaxAge = lang.Milliseconds(maxage).Duration()
		iClass.Namespace = ns
		iClass.SimpleName = alias
		iClass.Store = store
		iNS.Name = ns

		c.namespaces[ns] = iNS
		c.classes[fullname] = iClass
	}

	return g.Error()
}

func (inst *loader) loadStores(c *cache) error {
	const (
		prefix = "key-value-store."
		suffix = ".driver"
	)
	ac := inst.parent.AppContext
	g := ac.GetProperties().Getter().Required()
	ids := g.ListItems(prefix, suffix)

	for _, name := range ids {
		keyx := prefix + name + "."

		driver := g.GetString(keyx + "driver")
		username := g.GetString(keyx + "username")
		password := g.GetString(keyx + "password")
		host := g.GetString(keyx + "host")
		port := g.GetInt(keyx + "port")
		enabled := g.GetBool(keyx + "enabled")
		maxage1 := g.GetInt64(keyx + "default-max-age")

		maxage2 := lang.Milliseconds(maxage1)

		info := &keyvalues.StoreInfo{}
		info.Enabled = enabled
		info.Name = name
		info.Driver = driver
		info.Host = host
		info.Port = port
		info.Username = username
		info.Password = password
		info.MaxAge = maxage2.Duration()

		c.stores[name] = info
	}

	return g.Error()
}

func (inst *loader) loadDrivers(c *cache) error {
	reglist := inst.parent.RegList
	for _, r1 := range reglist {
		r2list := r1.ListRegistrations()
		for _, r2 := range r2list {
			info := r2.Driver
			name := info.Name
			c.drivers[name] = info
		}
	}
	return nil
}
