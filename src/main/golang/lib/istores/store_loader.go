package istores

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/keyvalues/src/main/golang/lib/info"
)

// DefaultStoreLoader ...
type DefaultStoreLoader struct {

	//starter:component

	_as func(keyvalues.StoreLoader) //starter:as("#")

	AppContext application.Context //starter:inject("context")
	Service    keyvalues.Service   //starter:inject("#")
	InfoGetter info.Getter         //starter:inject("#")
}

func (inst *DefaultStoreLoader) _impl() keyvalues.StoreLoader {
	return inst
}

// LoadStore ...
func (inst *DefaultStoreLoader) LoadStore(name string) (keyvalues.Store, error) {
	cfg, err := inst.getConfig(name)
	if err != nil {
		return nil, err
	}
	return inst.openStore(cfg)
}

func (inst *DefaultStoreLoader) info2config(info *keyvalues.StoreInfo) (*keyvalues.Configuration, error) {
	cfg := &keyvalues.Configuration{
		Name:          info.Name,
		Enabled:       info.Enabled,
		Driver:        info.Driver,
		Host:          info.Host,
		Port:          info.Port,
		Username:      info.Username,
		Password:      info.Password,
		DefaultMaxAge: info.MaxAge,
	}
	return cfg, nil
}

func (inst *DefaultStoreLoader) getConfig(name string) (*keyvalues.Configuration, error) {
	infolist := inst.InfoGetter.GetStoreInfoList()
	for _, info := range infolist {
		if info.Name == name {
			return inst.info2config(info)
		}
	}
	return nil, fmt.Errorf("cannot find store with name: '%s'", name)
}

func (inst *DefaultStoreLoader) openStore(cfg *keyvalues.Configuration) (keyvalues.Store, error) {
	ser := inst.Service
	driver, err := ser.GetDriver(cfg.Driver)
	if err != nil {
		return nil, err
	}
	return driver.Open(cfg)
}

////////////////////////////////////////////////////////////////////////////////
