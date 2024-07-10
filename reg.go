package keyvalues

import "time"

// DriverInfo ...
type DriverInfo struct {
	Name   string
	Driver Driver
}

// StoreInfo ...
type StoreInfo struct {
	Enabled bool
	Name    string

	Host     string
	Port     int
	Username string
	Password string
	Driver   string
	MaxAge   time.Duration

	Store Store
}

// NamespaceInfo ...
type NamespaceInfo struct {
	Name      NS
	Namespace Namespace
}

// ClassInfo ...
type ClassInfo struct {
	ID string // 在 properties 中的 ID

	Namespace  NS         // refer to NS.name
	Store      string     // refer to Store.name
	SimpleName SimpleName // simple name of this class
	MaxAge     time.Duration
	Enabled    bool

	Class  Class
	Loader ClassLoader
}

////////////////////////////////////////////////////////////////////////////////

// Registration 注册信息
type Registration struct {
	Enabled  bool
	Priority int

	Driver *DriverInfo

	// Store  *StoreInfo
	// NS     *NamespaceInfo
	// Class  *ClassInfo
}

// Registry ...
type Registry interface {
	ListRegistrations() []*Registration
}

////////////////////////////////////////////////////////////////////////////////
