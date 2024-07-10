package keyvalues

import "time"

// Options 表示 store 的操作选项
type Options struct {
	MaxAge time.Duration
}

// Configuration 表示 store 的配置
type Configuration struct {
	Name          string // name of this config
	Driver        string
	Host          string
	Port          int
	Username      string
	Password      string
	DefaultMaxAge time.Duration
	Enabled       bool
}

// Store 代表一个存储库
type Store interface {
	Put(key string, value []byte, opt *Options) error
	Get(key string) ([]byte, error)
	Contains(key string) (bool, error)
}

// StoreGetter ...
type StoreGetter interface {
	GetStore(name string) (Store, error)
}

// StoreLoader ...
type StoreLoader interface {
	LoadStore(name string) (Store, error)
}
