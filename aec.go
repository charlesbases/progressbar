package progressbar

import (
	"fmt"
	"io"

	"github.com/moby/term"
	"github.com/morikuni/aec"
)

// stdout .
func stdout() io.Writer {
	_, out, _ := term.StdStreams()
	return out
}

// clearline .
func clearline(out io.Writer) () {
	fmt.Fprint(out, aec.EraseLine(aec.EraseModes.All))
}

// cursorup .
func cursorup(out io.Writer, n uint) {
	if n > 0 {
		fmt.Fprint(out, aec.Up(n))
	}
}

// cursordown .
func cursordown(out io.Writer, n uint) {
	if n > 0 {
		fmt.Fprint(out, aec.Down(n))
	}
}
