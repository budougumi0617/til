package rate

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

// LimitConn is Limiter wrapper.
type LimitConn struct {
	lim *rate.Limiter
}

func (lc *LimitConn) wait(i int) {
	if err := lc.lim.Wait(context.Background()); err != nil {
		log.Printf("wait: %v\n", err)
		return
	}
	log.Printf("wait: done %d!\n", i)
}

func (lc *LimitConn) allow(i int) {
	if !lc.lim.Allow() {
		log.Printf("allow: cancel %d\n", i)
		return
	}
	log.Printf("allow: done %d!\n", i)
}

func (lc *LimitConn) reserve(i int) {
	r := lc.lim.Reserve()
	if i%2 == 0 {
		log.Printf("reserve: cancel %d!\n", i)
		r.Cancel()
	}
	if !r.OK() {
		// Not allowed to act! Did you remember to set lim.burst to be > 0 ?
		log.Printf("reserve: not ok %d!\n", i)
		return
	}
	time.Sleep(r.Delay())
	log.Printf("reserve: done %d!\n", i)
}

var (
	n = 10            // 発生するアクション数
	r = rate.Limit(2) // 1秒間に増加するトークン数
	m = 5             // 最大トークン数
)

// === RUN   TestWait
// 06:47:32 wait: done 9!
// 06:47:32 wait: done 3!
// 06:47:32 wait: done 2!
// 06:47:32 wait: done 6!
// 06:47:32 wait: done 4!
// 06:47:32 wait: done 5!
// 06:47:33 wait: done 1!
// 06:47:33 wait: done 0!
// 06:47:34 wait: done 8!
// 06:47:34 wait: done 7!
// --- PASS: TestWait (2.51s)
func TestWait(t *testing.T) {
	log.SetFlags(log.Ltime)

	lc := &LimitConn{
		lim: rate.NewLimiter(r, m),
	}
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lc.wait(i)
		}(i)
	}

	wg.Wait()
}

// === RUN   TestAllow
// 06:47:34 allow: done 0!
// 06:47:34 allow: done 1!
// 06:47:34 allow: done 2!
// 06:47:34 allow: cancel 5
// 06:47:34 allow: cancel 7
// 06:47:34 allow: done 9!
// 06:47:34 allow: done 3!
// 06:47:34 allow: cancel 4
// 06:47:34 allow: cancel 8
// 06:47:34 allow: cancel 6
// --- PASS: TestAllow (0.00s)
func TestAllow(t *testing.T) {
	log.SetFlags(log.Ltime)

	lc := &LimitConn{
		lim: rate.NewLimiter(r, m),
	}
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lc.allow(i)
		}(i)
	}

	wg.Wait()
}

// === RUN   TestReserve
// 06:47:34 reserve: done 9!
// 06:47:34 reserve: cancel 4!
// 06:47:34 reserve: done 4!
// 06:47:34 reserve: cancel 6!
// 06:47:34 reserve: done 6!
// 06:47:34 reserve: cancel 0!
// 06:47:34 reserve: done 0!
// 06:47:34 reserve: cancel 8!
// 06:47:34 reserve: cancel 2!
// 06:47:34 reserve: done 1!
// 06:47:35 reserve: done 5!
// 06:47:35 reserve: done 7!
// 06:47:36 reserve: done 8!
// 06:47:36 reserve: done 2!
// 06:47:37 reserve: done 3!
// --- PASS: TestReserve (2.51s)

func TestReserve(t *testing.T) {
	log.SetFlags(log.Ltime)

	lc := &LimitConn{
		lim: rate.NewLimiter(r, m),
	}
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lc.reserve(i)
		}(i)
	}

	wg.Wait()
}
