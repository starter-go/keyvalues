package keyvalues

// Driver 代表 store 的驱动
type Driver interface {
	Open(cfg *Configuration) (Store, error)
}

// DriverGetter ...
type DriverGetter interface {
	GetDriver(name string) (Driver, error)
}
