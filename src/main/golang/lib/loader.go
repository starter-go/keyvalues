package lib

// type holder struct {
// 	mutex sync.Mutex

// 	classes    map[keyvalues.ClassName]*keyvalues.ClassInfo
// 	namespaces map[keyvalues.NS]*keyvalues.NamespaceInfo
// 	stores     map[string]*keyvalues.StoreInfo
// 	drivers    map[string]*keyvalues.DriverInfo
// }

// func (inst *holder) getLocker() sync.Locker {
// 	return &inst.mutex
// }

// ////////////////////////////////////////////////////////////////////////////////

// type loader struct {
// 	classlist []*keyvalues.ClassInfo
// }

// func (inst *loader) load(src []keyvalues.Registry) (*holder, error) {

// 	h := new(holder)
// 	h.stores = make(map[string]*keyvalues.StoreInfo)
// 	h.drivers = make(map[string]*keyvalues.DriverInfo)
// 	h.classes = make(map[keyvalues.ClassName]*keyvalues.ClassInfo)
// 	h.namespaces = make(map[keyvalues.NS]*keyvalues.NamespaceInfo)

// 	for _, r1 := range src {
// 		err := inst.handleReg1(r1, h)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	err := inst.registerClasses(h)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return h, nil
// }

// func (inst *loader) handleReg1(r1 keyvalues.Registry, h *holder) error {

// 	if r1 == nil {
// 		return nil
// 	}

// 	r2list := r1.ListRegistrations()
// 	for _, r2 := range r2list {
// 		err := inst.handleReg2(r2, h)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (inst *loader) handleReg2(r2 *keyvalues.Registration, h *holder) error {

// 	if r2 == nil {
// 		return nil
// 	}

// 	if !r2.Enabled {
// 		return nil
// 	}

// 	err := inst.handleClass1(r2.Class, h)
// 	if err != nil {
// 		return err
// 	}

// 	err = inst.handleNS(r2.NS, h)
// 	if err != nil {
// 		return err
// 	}

// 	err = inst.handleStore(r2.Store, h)
// 	if err != nil {
// 		return err
// 	}

// 	err = inst.handleDriver(r2.Driver, h)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (inst *loader) handleDriver(item *keyvalues.DriverInfo, h *holder) error {
// 	if item == nil {
// 		return nil
// 	}
// 	if item.Driver == nil {
// 		return nil
// 	}
// 	name := item.Name
// 	h.drivers[name] = item
// 	return nil
// }

// func (inst *loader) handleStore(item *keyvalues.StoreInfo, h *holder) error {
// 	if item == nil {
// 		return nil
// 	}
// 	if item.Store == nil {
// 		return nil
// 	}
// 	name := item.Name
// 	h.stores[name] = item
// 	return nil
// }

// func (inst *loader) handleNS(item *keyvalues.NamespaceInfo, h *holder) error {
// 	if item == nil {
// 		return nil
// 	}
// 	if item.Namespace == nil {
// 		return nil
// 	}
// 	name := item.Name
// 	h.namespaces[name] = item
// 	return nil
// }

// func (inst *loader) handleClass1(item *keyvalues.ClassInfo, _ *holder) error {
// 	if item == nil {
// 		return nil
// 	}
// 	if item.Class == nil {
// 		return nil
// 	}
// 	inst.classlist = append(inst.classlist, item)
// 	return nil
// }

// func (inst *loader) handleClass2(item *keyvalues.ClassInfo, h *holder) error {

// 	nsName := item.Namespace
// 	ns := h.namespaces[nsName]
// 	if ns == nil {
// 		return fmt.Errorf("no key-value namespace with name '%s'", nsName)
// 	}

// 	alias := item.SimpleName
// 	cname := ns.Name.GetClassName(alias)
// 	h.classes[cname] = item
// 	return nil
// }

// func (inst *loader) registerClasses(h *holder) error {
// 	src := inst.classlist
// 	for _, item := range src {
// 		err := inst.handleClass2(item, h)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
