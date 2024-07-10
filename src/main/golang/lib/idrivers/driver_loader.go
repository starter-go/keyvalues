package idrivers

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/keyvalues/src/main/golang/lib/info"
)

// DriverLoader ...
type DriverLoader interface {
	Load(name string) (keyvalues.Driver, error)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultDriverLoader ...
type DefaultDriverLoader struct {

	//starter:component

	_as func(DriverLoader) //starter:as("#")

	AppContext application.Context //starter:inject("context")
	Service    keyvalues.Service   //starter:inject("#")
	InfoGetter info.Getter         //starter:inject("#")
}

func (inst *DefaultDriverLoader) _impl() DriverLoader {
	return inst
}

// Load ...
func (inst *DefaultDriverLoader) Load(name string) (keyvalues.Driver, error) {
	all := inst.InfoGetter.GetDriverInfoList()
	for _, item := range all {
		if item.Name == name {
			return item.Driver, nil
		}
	}
	return nil, fmt.Errorf("no key-value-store-driver with name '%s'", name)
}
