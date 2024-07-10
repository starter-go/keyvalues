package info

import "github.com/starter-go/keyvalues"

// Getter ...
type Getter interface {
	GetClassInfoList() []*keyvalues.ClassInfo
	GetNamespaceInfoList() []*keyvalues.NamespaceInfo
	GetStoreInfoList() []*keyvalues.StoreInfo
	GetDriverInfoList() []*keyvalues.DriverInfo
}
