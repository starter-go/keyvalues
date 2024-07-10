package inamespaces

import (
	"fmt"

	"github.com/starter-go/keyvalues"
	"github.com/starter-go/keyvalues/src/main/golang/lib/info"
)

// NamespaceLoader ...
type NamespaceLoader struct {

	//starter:component

	_as func(keyvalues.NamespaceLoader) //starter:as("#")

	InfoGetter info.Getter       //starter:inject("#")
	Service    keyvalues.Service //starter:inject("#")

}

func (inst *NamespaceLoader) _impl() keyvalues.NamespaceLoader {
	return inst
}

// LoadNamespace ...
func (inst *NamespaceLoader) LoadNamespace(ns keyvalues.NS) (keyvalues.Namespace, error) {

	all := inst.InfoGetter.GetNamespaceInfoList()
	for _, item := range all {
		if item.Name == ns {
			return inst.makeNamespace(item)
		}
	}
	return nil, fmt.Errorf("no namespace with name: '%s'", ns)
}

func (inst *NamespaceLoader) makeNamespace(item *keyvalues.NamespaceInfo) (keyvalues.Namespace, error) {
	ns := &nsImpl{
		service: inst.Service,
		info:    item,
	}
	item.Namespace = ns
	return ns, nil
}
