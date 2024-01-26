package progressbar

import (
	"bytes"
	"strconv"
)

// option .
type option interface {
	apply(r *Reader)
}

type formatFunc func(p *Progress) []byte

// apply .
func (f formatFunc) apply(r *Reader) () {
	r.format = f
}

// formatProgress .
func formatProgress(p *Progress, fn func(b *bytes.Buffer)) []byte {
	return bufferWriter(
		func(b *bytes.Buffer) []byte {
			if len(p.id) != 0 {
				b.WriteString(p.id + ": ")
			}

			fn(b)

			// (60/100)
			b.WriteString(" ")
			b.WriteString("(")
			b.WriteString(strconv.Itoa(int(p.current)))
			b.WriteString("/")
			b.WriteString(strconv.Itoa(int(p.size)))
			b.WriteString(")")

			b.WriteString(escape)
			return b.Bytes()
		},
	)
}

// Format .
func Format(fn func(p *Progress) []byte) formatFunc {
	return fn
}

// Symbol .
// style: [==============================>                   ]
func Symbol() formatFunc {
	return func(p *Progress) []byte {
		return formatProgress(
			p, func(b *bytes.Buffer) {
				percent := int(float64(p.current)/float64(p.size)*100) / 2

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
			},
		)
	}
}

// Percent .
// style: 60%
func Percent() formatFunc {
	return func(p *Progress) []byte {
		return formatProgress(
			p, func(b *bytes.Buffer) {
				percent := int(float64(p.current) / float64(p.size) * 100)

				b.WriteString(strconv.Itoa(percent) + "%")
			},
		)
	}
}
