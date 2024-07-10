package keyvalues

// NS 表示命名空间的字符串形式
type NS string

func (ns NS) String() string {
	return string(ns)
}

// GetClassName 取类型的全名
func (ns NS) GetClassName(alias ClassAlias) ClassName {
	s1 := ns.String()
	s2 := alias.String()
	str := s1 + "/" + s2
	return ClassName(str)
}

////////////////////////////////////////////////////////////////////////////////

// Namespace 代表命名空间
type Namespace interface {
	Name() NS

	GetClass(alias ClassAlias) (Class, error)
}

// NamespaceGetter ...
type NamespaceGetter interface {
	GetNamespace(ns NS) (Namespace, error)
}

// NamespaceLoader ...
type NamespaceLoader interface {
	LoadNamespace(ns NS) (Namespace, error)
}
