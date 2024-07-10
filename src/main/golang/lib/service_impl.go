package lib

import (
	"github.com/starter-go/keyvalues"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component

	_as func(keyvalues.Service) //starter:as("#")

	ClassGetter     keyvalues.ClassGetter     //starter:inject("#")
	NamespaceGetter keyvalues.NamespaceGetter //starter:inject("#")
	StoreGetter     keyvalues.StoreGetter     //starter:inject("#")
	DriverGetter    keyvalues.DriverGetter    //starter:inject("#")

}

func (inst *ServiceImpl) _impl() keyvalues.Service {
	return inst
}

// GetClass 。。。
func (inst *ServiceImpl) GetClass(name keyvalues.ClassName) (keyvalues.Class, error) {
	return inst.ClassGetter.GetClass(name)
}

// GetClassNS 。。。
func (inst *ServiceImpl) GetClassNS(ns keyvalues.NS, alias keyvalues.ClassAlias) (keyvalues.Class, error) {
	name := ns.GetClassName(alias)
	return inst.ClassGetter.GetClass(name)
}

// GetDriver 。。。
func (inst *ServiceImpl) GetDriver(name string) (keyvalues.Driver, error) {
	return inst.DriverGetter.GetDriver(name)
}

// GetNamespace 。。。
func (inst *ServiceImpl) GetNamespace(ns keyvalues.NS) (keyvalues.Namespace, error) {
	return inst.NamespaceGetter.GetNamespace(ns)
}

// GetStore 。。。
func (inst *ServiceImpl) GetStore(name string) (keyvalues.Store, error) {
	return inst.StoreGetter.GetStore(name)
}
