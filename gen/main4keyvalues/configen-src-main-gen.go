package main4keyvalues
import (
    p0ef6f2938 "github.com/starter-go/application"
    p21f95db42 "github.com/starter-go/keyvalues"
    p6fd3f5368 "github.com/starter-go/keyvalues/src/main/golang/lib"
    pa10248d56 "github.com/starter-go/keyvalues/src/main/golang/lib/iclasses"
    p68670769e "github.com/starter-go/keyvalues/src/main/golang/lib/idrivers"
    pedcc9550c "github.com/starter-go/keyvalues/src/main/golang/lib/inamespaces"
    p914cc6bb4 "github.com/starter-go/keyvalues/src/main/golang/lib/info"
    p2fd81db81 "github.com/starter-go/keyvalues/src/main/golang/lib/istores"
    p951b179c3 "github.com/starter-go/keyvalues/src/main/golang/lib/ram"
     "github.com/starter-go/application"
)

// type p6fd3f5368.ServiceImpl in package:github.com/starter-go/keyvalues/src/main/golang/lib
//
// id:com-6fd3f536816aba43-lib-ServiceImpl
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-Service
// scope:singleton
//
type p6fd3f53681_lib_ServiceImpl struct {
}

func (inst* p6fd3f53681_lib_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6fd3f536816aba43-lib-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6fd3f53681_lib_ServiceImpl) new() any {
    return &p6fd3f5368.ServiceImpl{}
}

func (inst* p6fd3f53681_lib_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6fd3f5368.ServiceImpl)
	nop(ie, com)

	
    com.ClassGetter = inst.getClassGetter(ie)
    com.NamespaceGetter = inst.getNamespaceGetter(ie)
    com.StoreGetter = inst.getStoreGetter(ie)
    com.DriverGetter = inst.getDriverGetter(ie)


    return nil
}


func (inst*p6fd3f53681_lib_ServiceImpl) getClassGetter(ie application.InjectionExt)p21f95db42.ClassGetter{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-ClassGetter").(p21f95db42.ClassGetter)
}


func (inst*p6fd3f53681_lib_ServiceImpl) getNamespaceGetter(ie application.InjectionExt)p21f95db42.NamespaceGetter{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-NamespaceGetter").(p21f95db42.NamespaceGetter)
}


func (inst*p6fd3f53681_lib_ServiceImpl) getStoreGetter(ie application.InjectionExt)p21f95db42.StoreGetter{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-StoreGetter").(p21f95db42.StoreGetter)
}


func (inst*p6fd3f53681_lib_ServiceImpl) getDriverGetter(ie application.InjectionExt)p21f95db42.DriverGetter{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-DriverGetter").(p21f95db42.DriverGetter)
}



// type pa10248d56.ClassCache in package:github.com/starter-go/keyvalues/src/main/golang/lib/iclasses
//
// id:com-a10248d56add93c1-iclasses-ClassCache
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-ClassGetter
// scope:singleton
//
type pa10248d56a_iclasses_ClassCache struct {
}

func (inst* pa10248d56a_iclasses_ClassCache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a10248d56add93c1-iclasses-ClassCache"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-ClassGetter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa10248d56a_iclasses_ClassCache) new() any {
    return &pa10248d56.ClassCache{}
}

func (inst* pa10248d56a_iclasses_ClassCache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa10248d56.ClassCache)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pa10248d56a_iclasses_ClassCache) getLoader(ie application.InjectionExt)p21f95db42.ClassLoader{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-ClassLoader").(p21f95db42.ClassLoader)
}



// type pa10248d56.DefaultClassLoader in package:github.com/starter-go/keyvalues/src/main/golang/lib/iclasses
//
// id:com-a10248d56add93c1-iclasses-DefaultClassLoader
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-ClassLoader
// scope:singleton
//
type pa10248d56a_iclasses_DefaultClassLoader struct {
}

func (inst* pa10248d56a_iclasses_DefaultClassLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a10248d56add93c1-iclasses-DefaultClassLoader"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-ClassLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa10248d56a_iclasses_DefaultClassLoader) new() any {
    return &pa10248d56.DefaultClassLoader{}
}

func (inst* pa10248d56a_iclasses_DefaultClassLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa10248d56.DefaultClassLoader)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Service = inst.getService(ie)
    com.Source = inst.getSource(ie)


    return nil
}


func (inst*pa10248d56a_iclasses_DefaultClassLoader) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pa10248d56a_iclasses_DefaultClassLoader) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}


func (inst*pa10248d56a_iclasses_DefaultClassLoader) getSource(ie application.InjectionExt)p914cc6bb4.Getter{
    return ie.GetComponent("#alias-914cc6bb44477f5c8662dc6c7044c04c-Getter").(p914cc6bb4.Getter)
}



// type p68670769e.DriverCache in package:github.com/starter-go/keyvalues/src/main/golang/lib/idrivers
//
// id:com-68670769ee8944ea-idrivers-DriverCache
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-DriverGetter
// scope:singleton
//
type p68670769ee_idrivers_DriverCache struct {
}

func (inst* p68670769ee_idrivers_DriverCache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-68670769ee8944ea-idrivers-DriverCache"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-DriverGetter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p68670769ee_idrivers_DriverCache) new() any {
    return &p68670769e.DriverCache{}
}

func (inst* p68670769ee_idrivers_DriverCache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p68670769e.DriverCache)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*p68670769ee_idrivers_DriverCache) getLoader(ie application.InjectionExt)p68670769e.DriverLoader{
    return ie.GetComponent("#alias-68670769ee8944ea68a90acd02542633-DriverLoader").(p68670769e.DriverLoader)
}



// type p68670769e.DefaultDriverLoader in package:github.com/starter-go/keyvalues/src/main/golang/lib/idrivers
//
// id:com-68670769ee8944ea-idrivers-DefaultDriverLoader
// class:
// alias:alias-68670769ee8944ea68a90acd02542633-DriverLoader
// scope:singleton
//
type p68670769ee_idrivers_DefaultDriverLoader struct {
}

func (inst* p68670769ee_idrivers_DefaultDriverLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-68670769ee8944ea-idrivers-DefaultDriverLoader"
	r.Classes = ""
	r.Aliases = "alias-68670769ee8944ea68a90acd02542633-DriverLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p68670769ee_idrivers_DefaultDriverLoader) new() any {
    return &p68670769e.DefaultDriverLoader{}
}

func (inst* p68670769ee_idrivers_DefaultDriverLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p68670769e.DefaultDriverLoader)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Service = inst.getService(ie)
    com.InfoGetter = inst.getInfoGetter(ie)


    return nil
}


func (inst*p68670769ee_idrivers_DefaultDriverLoader) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p68670769ee_idrivers_DefaultDriverLoader) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}


func (inst*p68670769ee_idrivers_DefaultDriverLoader) getInfoGetter(ie application.InjectionExt)p914cc6bb4.Getter{
    return ie.GetComponent("#alias-914cc6bb44477f5c8662dc6c7044c04c-Getter").(p914cc6bb4.Getter)
}



// type pedcc9550c.NamespaceCache in package:github.com/starter-go/keyvalues/src/main/golang/lib/inamespaces
//
// id:com-edcc9550ce15772a-inamespaces-NamespaceCache
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-NamespaceGetter
// scope:singleton
//
type pedcc9550ce_inamespaces_NamespaceCache struct {
}

func (inst* pedcc9550ce_inamespaces_NamespaceCache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-edcc9550ce15772a-inamespaces-NamespaceCache"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-NamespaceGetter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pedcc9550ce_inamespaces_NamespaceCache) new() any {
    return &pedcc9550c.NamespaceCache{}
}

func (inst* pedcc9550ce_inamespaces_NamespaceCache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pedcc9550c.NamespaceCache)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pedcc9550ce_inamespaces_NamespaceCache) getLoader(ie application.InjectionExt)p21f95db42.NamespaceLoader{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-NamespaceLoader").(p21f95db42.NamespaceLoader)
}



// type pedcc9550c.NamespaceLoader in package:github.com/starter-go/keyvalues/src/main/golang/lib/inamespaces
//
// id:com-edcc9550ce15772a-inamespaces-NamespaceLoader
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-NamespaceLoader
// scope:singleton
//
type pedcc9550ce_inamespaces_NamespaceLoader struct {
}

func (inst* pedcc9550ce_inamespaces_NamespaceLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-edcc9550ce15772a-inamespaces-NamespaceLoader"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-NamespaceLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pedcc9550ce_inamespaces_NamespaceLoader) new() any {
    return &pedcc9550c.NamespaceLoader{}
}

func (inst* pedcc9550ce_inamespaces_NamespaceLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pedcc9550c.NamespaceLoader)
	nop(ie, com)

	
    com.InfoGetter = inst.getInfoGetter(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pedcc9550ce_inamespaces_NamespaceLoader) getInfoGetter(ie application.InjectionExt)p914cc6bb4.Getter{
    return ie.GetComponent("#alias-914cc6bb44477f5c8662dc6c7044c04c-Getter").(p914cc6bb4.Getter)
}


func (inst*pedcc9550ce_inamespaces_NamespaceLoader) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}



// type p914cc6bb4.GetterImpl in package:github.com/starter-go/keyvalues/src/main/golang/lib/info
//
// id:com-914cc6bb44477f5c-info-GetterImpl
// class:
// alias:alias-914cc6bb44477f5c8662dc6c7044c04c-Getter
// scope:singleton
//
type p914cc6bb44_info_GetterImpl struct {
}

func (inst* p914cc6bb44_info_GetterImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-914cc6bb44477f5c-info-GetterImpl"
	r.Classes = ""
	r.Aliases = "alias-914cc6bb44477f5c8662dc6c7044c04c-Getter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p914cc6bb44_info_GetterImpl) new() any {
    return &p914cc6bb4.GetterImpl{}
}

func (inst* p914cc6bb44_info_GetterImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p914cc6bb4.GetterImpl)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.RegList = inst.getRegList(ie)


    return nil
}


func (inst*p914cc6bb44_info_GetterImpl) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p914cc6bb44_info_GetterImpl) getRegList(ie application.InjectionExt)[]p21f95db42.Registry{
    dst := make([]p21f95db42.Registry, 0)
    src := ie.ListComponents(".class-21f95db421796c61fc702c5dfd6515de-Registry")
    for _, item1 := range src {
        item2 := item1.(p21f95db42.Registry)
        dst = append(dst, item2)
    }
    return dst
}



// type p2fd81db81.StoreCache in package:github.com/starter-go/keyvalues/src/main/golang/lib/istores
//
// id:com-2fd81db8107fb5be-istores-StoreCache
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-StoreGetter
// scope:singleton
//
type p2fd81db810_istores_StoreCache struct {
}

func (inst* p2fd81db810_istores_StoreCache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2fd81db8107fb5be-istores-StoreCache"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-StoreGetter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2fd81db810_istores_StoreCache) new() any {
    return &p2fd81db81.StoreCache{}
}

func (inst* p2fd81db810_istores_StoreCache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2fd81db81.StoreCache)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*p2fd81db810_istores_StoreCache) getLoader(ie application.InjectionExt)p21f95db42.StoreLoader{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-StoreLoader").(p21f95db42.StoreLoader)
}



// type p2fd81db81.DefaultStoreLoader in package:github.com/starter-go/keyvalues/src/main/golang/lib/istores
//
// id:com-2fd81db8107fb5be-istores-DefaultStoreLoader
// class:
// alias:alias-21f95db421796c61fc702c5dfd6515de-StoreLoader
// scope:singleton
//
type p2fd81db810_istores_DefaultStoreLoader struct {
}

func (inst* p2fd81db810_istores_DefaultStoreLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2fd81db8107fb5be-istores-DefaultStoreLoader"
	r.Classes = ""
	r.Aliases = "alias-21f95db421796c61fc702c5dfd6515de-StoreLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2fd81db810_istores_DefaultStoreLoader) new() any {
    return &p2fd81db81.DefaultStoreLoader{}
}

func (inst* p2fd81db810_istores_DefaultStoreLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2fd81db81.DefaultStoreLoader)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Service = inst.getService(ie)
    com.InfoGetter = inst.getInfoGetter(ie)


    return nil
}


func (inst*p2fd81db810_istores_DefaultStoreLoader) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p2fd81db810_istores_DefaultStoreLoader) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}


func (inst*p2fd81db810_istores_DefaultStoreLoader) getInfoGetter(ie application.InjectionExt)p914cc6bb4.Getter{
    return ie.GetComponent("#alias-914cc6bb44477f5c8662dc6c7044c04c-Getter").(p914cc6bb4.Getter)
}



// type p951b179c3.MemoryDriver in package:github.com/starter-go/keyvalues/src/main/golang/lib/ram
//
// id:com-951b179c39347766-ram-MemoryDriver
// class:class-21f95db421796c61fc702c5dfd6515de-Registry
// alias:
// scope:singleton
//
type p951b179c39_ram_MemoryDriver struct {
}

func (inst* p951b179c39_ram_MemoryDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-951b179c39347766-ram-MemoryDriver"
	r.Classes = "class-21f95db421796c61fc702c5dfd6515de-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p951b179c39_ram_MemoryDriver) new() any {
    return &p951b179c3.MemoryDriver{}
}

func (inst* p951b179c39_ram_MemoryDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p951b179c3.MemoryDriver)
	nop(ie, com)

	


    return nil
}


