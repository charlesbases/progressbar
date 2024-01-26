package progressbar

const escape = "\r"

// Progress .
type Progress struct {
	id string

	size    uint
	current uint

	// chan of Reader
	c chan *Progress
}

// freed .
func (p *Progress) freed() () {
	p.c = nil
	p.id = ""
	p.size = 0
	p.current = 0
	progressPool.Put(p)
}

// Incr 增加当前进度
func (p *Progress) Incr(n uint) () {
	if p.size > 0 && p.current < p.size {
		p.current += n
		if p.current > p.size {
			p.current = p.size
		}

		p.c <- newProgress(
			func(pp *Progress) {
				pp.id = p.id
				pp.size = p.size
				pp.current = p.current
			},
		)
	}
}
