package inamespaces

import "github.com/starter-go/keyvalues"

// nsImpl ...
type nsImpl struct {
	service keyvalues.Service
	info    *keyvalues.NamespaceInfo
}

func (inst *nsImpl) _impl() keyvalues.Namespace {
	return inst
}

// ListRegistrations ...
func (inst *nsImpl) Name() keyvalues.NS {
	return inst.info.Name
}

// ListRegistrations ...
func (inst *nsImpl) GetClass(alias keyvalues.ClassAlias) (keyvalues.Class, error) {
	ns := inst.Name()
	cn := ns.GetClassName(alias)
	return inst.service.GetClass(cn)
}
