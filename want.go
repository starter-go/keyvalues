package keyvalues

// Want 表示一个查询条件
type Want struct {
	Namespace string
	Class     SimpleName
	ID        string
}
