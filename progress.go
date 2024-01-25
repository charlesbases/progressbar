package progressbar

import (
	"bytes"
	"strconv"
	"sync"
)

const escape = "\r"

var buffPool sync.Pool

// Progress .
type Progress struct {
	id string

	size    uint
	current uint

	// chan of Reader
	c chan *Progress
}

// newBuffer .
func newBuffer() *bytes.Buffer {
	if v := buffPool.Get(); v != nil {
		b := v.(*bytes.Buffer)
		return b
	}
	return bytes.NewBuffer(nil)
}

// format .
func (p *Progress) format() []byte {
	var b = newBuffer()
	defer func() {
		b.Reset()
		buffPool.Put(b)
	}()

	if len(p.id) != 0 {
		b.WriteString(p.id + ": ")
	}

	p.formatProgress(b)

	return b.Bytes()
}

// formatProgress .
func (p *Progress) formatProgress(b *bytes.Buffer) {
	// p.formatProgressSymbol(b)
	p.formatProgressPercent(b)
	b.WriteString(escape)
}

// formatProgressSymbol .
// style: [==============================>                   ] (60/100)
func (p *Progress) formatProgressSymbol(b *bytes.Buffer) () {
	percent := int(float64(p.current)/float64(p.size)*100) / 2

	{
		b.WriteString("[")
		for i := 0; i < percent; i++ {
			b.WriteString("=")
		}
		if p.current < p.size {
			b.WriteString(">")
		}
		for i := 0; i < 50-percent-1; i++ {
			b.WriteString(" ")
		}
		b.WriteString("]")
	}

	{
		b.WriteString(" ")
		b.WriteString("(")
		b.WriteString(strconv.Itoa(int(p.current)))
		b.WriteString("/")
		b.WriteString(strconv.Itoa(int(p.size)))
		b.WriteString(")")
	}
}

// formatProgressPercent .
// style: 60% (60/100)
func (p *Progress) formatProgressPercent(b *bytes.Buffer) () {
	percent := int(float64(p.current) / float64(p.size) * 100)
	b.WriteString(strconv.Itoa(percent) + "%")

	{
		b.WriteString(" ")
		b.WriteString("(")
		b.WriteString(strconv.Itoa(int(p.current)))
		b.WriteString("/")
		b.WriteString(strconv.Itoa(int(p.size)))
		b.WriteString(")")
	}
}

// Increment .
func (p *Progress) Increment(n uint) () {
	if p.size > 0 && p.current < p.size {
		p.current += n
		if p.current > p.size {
			p.current = p.size
		}

		p.c <- &Progress{
			id:      p.id,
			size:    p.size,
			current: p.current,
		}
	}
}
