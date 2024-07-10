package keyvalues

// ClassName 表示类型的全名 （namespace + simpleName）
type ClassName string

// ClassAlias 表示类型的别名
type ClassAlias string

// SimpleName 表示类型的短名称
type SimpleName = ClassAlias

// Class 代表一个类型
type Class interface {
	Name() ClassName

	Alias() ClassAlias

	Namespace() Namespace

	// GetEntry(id string) Entry

	GetTextEntry(id string) TextEntry

	GetBinaryEntry(id string) BinaryEntry

	GetJSONEntry(id string) JSONEntry
}

// ClassLoader ...
type ClassLoader interface {
	LoadClass(name ClassName) (Class, error)
}

// ClassGetter ...
type ClassGetter interface {
	GetClass(name ClassName) (Class, error)
}

////////////////////////////////////////////////////////////////////////////////

func (cn ClassName) String() string {
	return string(cn)
}

////////////////////////////////////////////////////////////////////////////////

func (cn ClassAlias) String() string {
	return string(cn)
}

////////////////////////////////////////////////////////////////////////////////
