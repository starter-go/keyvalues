package unit

import (
	"fmt"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues"
)

// DemoUnit ... 单元测试示例
type DemoUnit struct {

	//starter:component

	// _as func(units.Units) //starter:as(".")

	Service keyvalues.Service //starter:inject("#")

}

// func (inst *DemoUnit) _impl() units.Units { return inst }

// Life ...
func (inst *DemoUnit) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.test1,
	}
}

// Units ...
func (inst *DemoUnit) test1() error {

	ser := inst.Service
	ns, err := ser.GetNamespace("kvs.ns.ns1")
	if err != nil {
		return err
	}

	cl, err := ns.GetClass("Class1")
	if err != nil {
		return err
	}

	value1 := "hello, kvs"
	ent := cl.GetTextEntry("demo1")
	err = ent.Put(value1, nil)
	if err != nil {
		return err
	}

	time.Sleep(time.Second)

	value2, err := ent.Get()
	if err != nil {
		return err
	}

	if value1 != value2 {
		return fmt.Errorf("value1 != value2")
	}

	return nil
}
