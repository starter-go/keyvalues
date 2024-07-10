package keyvalues

// Service 提供面向客户端的服务
type Service interface {

	// GetDriver(name string) (Driver, error)
	// GetStore(name string) (Store, error)
	// GetNamespace(ns NS) (Namespace, error)
	// GetClass(cname ClassName) (Class, error)

	GetClassNS(ns NS, alias ClassAlias) (Class, error)

	ClassGetter
	DriverGetter
	NamespaceGetter
	StoreGetter
}
