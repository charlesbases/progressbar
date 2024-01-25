package progressbar

import (
	"sync"
	"testing"
	"time"
)

// _progressbar .
func _progressbar() () {
	r := NewReader()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		var size uint = 100
		p := r.NewProgress("progressbar1", size)

		var i int
		for i < int(size) {
			p.Increment(1)
			<-time.After(30 * time.Millisecond)
			i++
		}

		wg.Done()
	}()

	go func() {
		var size uint = 1000
		p := r.NewProgress("progressbar2", size)

		var i int
		for i < int(size) {
			p.Increment(1)
			<-time.After(6 * time.Millisecond)
			i++
		}

		wg.Done()
	}()

	wg.Wait()
	r.Close()
}
func TestProgressbar(t *testing.T) {
	_progressbar()
}

// BenchmarkProgressbar-16                1        6555312100 ns/op          570752 B/op      17462 allocs/op
func BenchmarkProgressbar(b *testing.B) {
	var bench = func(f func()) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			f()
		}
		b.StopTimer()
	}

	bench(
		func() {
			_progressbar()
		},
	)

	// fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}
