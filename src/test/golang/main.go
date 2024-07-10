package main

import (
	"os"

	"github.com/starter-go/keyvalues/modules/keyvalues"
	"github.com/starter-go/starter"
)

func main() {
	m := keyvalues.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
