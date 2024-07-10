package iclasses

import (
	"encoding/json"

	"github.com/starter-go/keyvalues"
)

type myClassImpl struct {
	info     *keyvalues.ClassInfo
	fullname keyvalues.ClassName
	ns       keyvalues.Namespace
	store    keyvalues.Store
}

func (inst *myClassImpl) _impl() keyvalues.Class {
	return inst
}

func (inst *myClassImpl) Name() keyvalues.ClassName {
	return inst.fullname
}

func (inst *myClassImpl) Alias() keyvalues.ClassAlias {
	return inst.info.SimpleName
}

func (inst *myClassImpl) Namespace() keyvalues.Namespace {
	return inst.ns
}

func (inst *myClassImpl) GetTextEntry(id string) keyvalues.TextEntry {
	return &myTextEntry{parent: inst, id: id}
}

func (inst *myClassImpl) GetBinaryEntry(id string) keyvalues.BinaryEntry {
	return &myBinaryEntry{parent: inst, id: id}
}

func (inst *myClassImpl) GetJSONEntry(id string) keyvalues.JSONEntry {
	return &myJSONEntry{parent: inst, id: id}
}

func (inst *myClassImpl) keyOf(id string) string {
	clname := inst.fullname
	return clname.String() + "#" + id
}

func (inst *myClassImpl) fetch(id string) ([]byte, error) {
	key := inst.keyOf(id)
	return inst.store.Get(key)
}

func (inst *myClassImpl) put(id string, value []byte, opt *keyvalues.Options) error {

	if opt == nil {
		opt = new(keyvalues.Options)
	}
	if opt.MaxAge < 1 {
		opt.MaxAge = inst.info.MaxAge
	}

	key := inst.keyOf(id)
	return inst.store.Put(key, value, opt)
}

////////////////////////////////////////////////////////////////////////////////

type myBinaryEntry struct {
	parent *myClassImpl
	id     string
}

func (inst *myBinaryEntry) _impl() keyvalues.BinaryEntry {
	return inst
}

func (inst *myBinaryEntry) Class() keyvalues.Class {
	return inst.parent
}

func (inst *myBinaryEntry) Key() string {
	return inst.id
}

func (inst *myBinaryEntry) Put(value []byte, opt *keyvalues.Options) error {
	id := inst.id
	// store := inst.parent.store
	// return store.Put(id, value, opt)
	return inst.parent.put(id, value, opt)
}

func (inst *myBinaryEntry) Get() ([]byte, error) {
	id := inst.id
	// store := inst.parent.store
	return inst.parent.fetch(id) // store.Get(id)
}

////////////////////////////////////////////////////////////////////////////////

type myTextEntry struct {
	id     string
	parent *myClassImpl
}

func (inst *myTextEntry) _impl() keyvalues.TextEntry {
	return inst
}

func (inst *myTextEntry) Class() keyvalues.Class {
	return inst.parent
}

func (inst *myTextEntry) Key() string {
	return inst.id
}

func (inst *myTextEntry) Put(value string, opt *keyvalues.Options) error {
	id := inst.id
	// store := inst.parent.store
	data := []byte(value)
	return inst.parent.put(id, data, opt)
}

func (inst *myTextEntry) Get() (string, error) {
	id := inst.id
	// store := inst.parent.store
	data, err := inst.parent.fetch(id) // store.Get(id)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

////////////////////////////////////////////////////////////////////////////////

type myJSONEntry struct {
	id     string
	parent *myClassImpl
}

func (inst *myJSONEntry) _impl() keyvalues.JSONEntry {
	return inst
}

func (inst *myJSONEntry) Class() keyvalues.Class {
	return inst.parent
}

func (inst *myJSONEntry) Key() string {
	return inst.id
}

func (inst *myJSONEntry) Put(value any, opt *keyvalues.Options) error {
	id := inst.id
	// store := inst.parent.store
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return inst.parent.put(id, data, opt)
}

func (inst *myJSONEntry) Get(value any) error {
	id := inst.id
	// store := inst.parent.store
	data, err := inst.parent.fetch(id) // store.Get(id)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}

////////////////////////////////////////////////////////////////////////////////
