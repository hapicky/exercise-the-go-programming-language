package areader

import (
	"errors"
	"path/filepath"
	"sync"
	"sync/atomic"
)

// アーカイブ形式の名称と対応関数
type format struct {
	name      string
	filenames func(path string) []string
}

// 対応形式
var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

// 対応形式を登録する
func RegisterFormat(name string, filenames func(path string) []string) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, filenames}))
	formatsMu.Unlock()
}

// アーカイブのパスから形式を特定します
func detect(path string) format {
	formats, _ := atomicFormats.Load().([]format)
	ext := filepath.Ext(path)
	for _, f := range formats {
		if ext == ("." + f.name) {
			return f
		}
	}
	return format{}
}

var ErrFormat = errors.New("areader: unknown format")

func Filenames(name string) ([]string, error) {
	format := detect(name)
	if format.filenames == nil {
		return nil, ErrFormat
	}
	return format.filenames(name), nil
}
