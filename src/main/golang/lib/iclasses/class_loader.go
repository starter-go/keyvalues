package iclasses

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/keyvalues/src/main/golang/lib/info"
)

// DefaultClassLoader ...
type DefaultClassLoader struct {

	//starter:component

	_as func(keyvalues.ClassLoader) //starter:as("#")

	AppContext application.Context //starter:inject("context")
	Service    keyvalues.Service   //starter:inject("#")
	Source     info.Getter         //starter:inject("#")

}

func (inst *DefaultClassLoader) _impl() keyvalues.ClassLoader {
	return inst
}

// LoadClass ...
func (inst *DefaultClassLoader) LoadClass(want keyvalues.ClassName) (keyvalues.Class, error) {
	all := inst.Source.GetClassInfoList()
	for _, item := range all {
		ns := item.Namespace
		alias := item.SimpleName
		have := ns.GetClassName(alias)
		if have == want {
			return inst.makeResult(have, item)
		}
	}
	return nil, fmt.Errorf("no class with name: '%s'", want)
}

func (inst *DefaultClassLoader) makeResult(name keyvalues.ClassName, item *keyvalues.ClassInfo) (keyvalues.Class, error) {

	ns, err := inst.Service.GetNamespace(item.Namespace)
	if err != nil {
		return nil, err
	}

	store, err := inst.Service.GetStore(item.Store)
	if err != nil {
		return nil, err
	}

	cl := &myClassImpl{
		fullname: name,
		info:     item,
		ns:       ns,
		store:    store,
	}
	return cl, nil
}

////////////////////////////////////////////////////////////////////////////////

// // type myClassGroup struct{}

// ////////////////////////////////////////////////////////////////////////////////

// type innerDefaultClassLoader struct {

// 	// groups  []*myClassGroup

// 	parent  *DefaultClassLoader
// 	classes []*keyvalues.ClassInfo
// }

// func (inst *innerDefaultClassLoader) load() error {
// 	return inst.loadGroups1()
// }

// func (inst *innerDefaultClassLoader) results() []*keyvalues.Registration {
// 	src := inst.classes
// 	dst := make([]*keyvalues.Registration, 0)
// 	for _, cl := range src {
// 		reg := &keyvalues.Registration{
// 			Enabled:  cl.Enabled,
// 			Priority: 0,
// 			Class:    cl,
// 		}
// 		dst = append(dst, reg)
// 	}
// 	return dst
// }

// func (inst *innerDefaultClassLoader) loadGroups1() error {
// 	const (
// 		prefix = "key-value-class-set."
// 		suffix = ".properties"
// 	)
// 	ac := inst.parent.AppContext
// 	ptable := ac.GetProperties()
// 	getter := ptable.Getter()
// 	ids := getter.ListItems(prefix, suffix)
// 	for _, id := range ids {
// 		ref, err := ptable.GetPropertyRequired(prefix + id + suffix)
// 		if err != nil {
// 			return err
// 		}
// 		err = inst.loadGroups2(ref)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (inst *innerDefaultClassLoader) loadGroups2(ref string) error {
// 	const (
// 		prefix = "key-value-class."
// 		suffix = ".name"
// 	)
// 	ac := inst.parent.AppContext
// 	text, err := ac.GetResources().ReadText(ref)
// 	if err != nil {
// 		return err
// 	}

// 	ptable, err := properties.Parse(text, nil)
// 	if err != nil {
// 		return err
// 	}

// 	getter := ptable.Getter()
// 	ids := getter.ListItems(prefix, suffix)
// 	for _, id := range ids {
// 		err = inst.loadClass(prefix, id, getter)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (inst *innerDefaultClassLoader) loadClass(prefix1, id string, getter properties.Getter) error {

// 	prefix2 := prefix1 + id + "."
// 	info := &keyvalues.ClassInfo{}
// 	cl := &myClassImpl{}
// 	cl.info = info

// 	name := getter.GetString(prefix2 + "name")
// 	store := getter.GetString(prefix2 + "store")
// 	ns := getter.GetString(prefix2 + "namespace")
// 	enabled := getter.GetBool(prefix2 + "enabled")
// 	maxage := getter.GetInt64(prefix2 + "max-age")

// 	maxage2 := lang.Milliseconds(maxage)
// 	err := getter.Error()
// 	if err != nil {
// 		return err
// 	}

// 	info.ID = id
// 	info.Class = cl
// 	info.MaxAge = maxage2.Duration()
// 	info.Namespace = keyvalues.NS(ns)
// 	info.Store = store
// 	info.SimpleName = keyvalues.ClassAlias(name)
// 	info.Enabled = enabled

// 	inst.classes = append(inst.classes, info)
// 	return nil
// }
