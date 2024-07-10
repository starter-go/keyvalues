package keyvalues

import (
	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/keyvalues/gen/main4keyvalues"
	"github.com/starter-go/keyvalues/gen/test4keyvalues"
	"github.com/starter-go/starter"
)

// Module  ...
func Module() application.Module {
	mb := keyvalues.NewMainModule()
	mb.Components(main4keyvalues.ExportComponents)
	mb.Depend(starter.Module())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := keyvalues.NewTestModule()
	mb.Components(test4keyvalues.ExportComponents)
	mb.Depend(Module())
	return mb.Create()
}
