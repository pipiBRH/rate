package rate

import (
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           // request limit
	begin time.Time     // windows start time
	cycle time.Duration // time window
	count int           // reueests in one window
	lock  sync.Mutex
}

func NewLimiter(r int, cycle time.Duration) *LimitRate {
	return &LimitRate{
		rate:  r,
		begin: time.Now(),
		cycle: cycle,
		count: 0,
	}
}

func (l *LimitRate) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now()
	if now.Sub(l.begin) >= l.cycle {
		l.reset(now)
		return true
	}

	if l.count == l.rate {
		return false
	}

	l.count++
	return true
}

func (l *LimitRate) reset(t time.Time) {
	l.begin = t
	l.count = 1
}

func (l *LimitRate) GetCount() int {
	return l.count
}
