package progressbar

import (
	"fmt"
	"io"
)

// Reader .
type Reader struct {
	// input
	c chan *Progress

	// ids for Progress
	ids map[string]int

	output io.Writer
}

// NewReader .
// newLine: 是否从新行打印进度条
func NewReader() *Reader {
	r := &Reader{
		c:      make(chan *Progress, 1),
		ids:    make(map[string]int, 0),
		output: stdout(),
	}

	go r.daemon()
	return r
}

// daemon .
func (r *Reader) daemon() () {
	for {
		select {
		case p, ok := <-r.c:
			if ok {
				r.display(p)
			} else {
				return
			}
		}
	}
}

// display .
func (r *Reader) display(p *Progress) () {
	var diff uint

	// 判断当前 Progress 是否打印过
	{
		line, ok := r.ids[p.id]
		if !ok {
			line = len(r.ids)
			r.ids[p.id] = line
			fmt.Fprint(r.output, "\n")
		}
		diff = uint(len(r.ids) - line)

		// 移动光标到当前行
		cursorup(r.output, diff)

		// 将光标放置在行首
		clearline(r.output)
	}

	// 写入 Progress
	r.output.Write(p.format())

	// 将光标重置到最后一行
	cursordown(r.output, diff)
}

// NewProgress .
func (r *Reader) NewProgress(id string, size uint) *Progress {
	return &Progress{id: id, size: size, c: r.c}
}

// Close .
func (r *Reader) Close() () {
	close(r.c)
	r = nil
}
