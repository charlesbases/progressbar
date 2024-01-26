package progressbar

import (
	"sync"
	"testing"
	"time"
)

// pb .
// output: Percent
// progressbar1: 100% (10/10)
// progressbar2: 100% (100/100)
// progressbar3: 100% (1000/1000)
//
// output: Symbol
// progressbar1: [==================================================] (10/10)
// progressbar2: [==================================================] (100/100)
// progressbar3: [==================================================] (1000/1000)
func pb() () {
	r := NewReader(Symbol())
	// r := NewReader(
	// 	Format(
	// 		func(p *Progress) []byte {
	// 			return nil
	// 		},
	// 	),
	// )

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		var size uint = 10
		p := r.NewProgress("progressbar1", size)

		var i int
		for i < int(size) {
			p.Incr(1)
			<-time.After(500 * time.Millisecond)
			i++
		}

		wg.Done()
	}()

	go func() {
		var size uint = 100
		p := r.NewProgress("progressbar2", size)

		var i int
		for i < int(size) {
			p.Incr(1)
			<-time.After(30 * time.Millisecond)
			i++
		}

		wg.Done()
	}()

	go func() {
		var size uint = 1000
		p := r.NewProgress("progressbar3", size)

		var i int
		for i < int(size) {
			p.Incr(1)
			<-time.After(6 * time.Millisecond)
			i++
		}

		wg.Done()
	}()

	wg.Wait()
	r.Close()
}

func TestProgressbar(t *testing.T) {
	pb()
}

// BenchmarkProgressbar-16                1        6557205100 ns/op          419128 B/op      16515 allocs/op
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
			pb()
		},
	)
}
