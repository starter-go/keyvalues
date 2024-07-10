package test4keyvalues
import (
    p21f95db42 "github.com/starter-go/keyvalues"
    p4e89af650 "github.com/starter-go/keyvalues/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p4e89af650.DemoUnit in package:github.com/starter-go/keyvalues/src/test/golang/unit
//
// id:com-4e89af65070d9b26-unit-DemoUnit
// class:
// alias:
// scope:singleton
//
type p4e89af6507_unit_DemoUnit struct {
}

func (inst* p4e89af6507_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4e89af65070d9b26-unit-DemoUnit"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4e89af6507_unit_DemoUnit) new() any {
    return &p4e89af650.DemoUnit{}
}

func (inst* p4e89af6507_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4e89af650.DemoUnit)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p4e89af6507_unit_DemoUnit) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}


