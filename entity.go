package keyvalues

// Entry 代表 store 中的一个条目
type Entry interface {
	Class() Class

	Key() string
}

// TextEntry 是 文本 形式的 Entry
type TextEntry interface {
	Entry

	Put(value string, opt *Options) error

	Get() (string, error)
}

// BinaryEntry 是 二进制 形式的 Entry
type BinaryEntry interface {
	Entry

	Put(value []byte, opt *Options) error

	Get() ([]byte, error)
}

// JSONEntry 是 JSON 形式的 Entry
type JSONEntry interface {
	Entry

	Put(value any, opt *Options) error

	Get(value any) error
}
