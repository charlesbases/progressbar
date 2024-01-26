package progressbar

import (
	"bytes"
	"sync"
)

var (
	bufferPool   sync.Pool
	progressPool sync.Pool
)

// bufferWriter .
func bufferWriter(f func(b *bytes.Buffer) []byte) []byte {
	var b *bytes.Buffer
	if v := bufferPool.Get(); v != nil {
		b = v.(*bytes.Buffer)
	} else {
		b = bytes.NewBuffer(nil)
	}

	data := f(b)
	b.Reset()
	bufferPool.Put(b)
	return data
}

// newProgress .
func newProgress(f func(p *Progress)) *Progress {
	var p *Progress
	if v := progressPool.Get(); v != nil {
		p = v.(*Progress)
	} else {
		p = new(Progress)
	}

	f(p)
	return p
}
