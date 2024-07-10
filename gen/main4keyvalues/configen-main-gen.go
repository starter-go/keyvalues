package main4keyvalues

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p2fd81db810_istores_DefaultStoreLoader{})
    inst.register(&p2fd81db810_istores_StoreCache{})
    inst.register(&p68670769ee_idrivers_DefaultDriverLoader{})
    inst.register(&p68670769ee_idrivers_DriverCache{})
    inst.register(&p6fd3f53681_lib_ServiceImpl{})
    inst.register(&p914cc6bb44_info_GetterImpl{})
    inst.register(&p951b179c39_ram_MemoryDriver{})
    inst.register(&pa10248d56a_iclasses_ClassCache{})
    inst.register(&pa10248d56a_iclasses_DefaultClassLoader{})
    inst.register(&pedcc9550ce_inamespaces_NamespaceCache{})
    inst.register(&pedcc9550ce_inamespaces_NamespaceLoader{})


    return nil
}
