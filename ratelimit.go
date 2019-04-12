package rate

import (
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           //請求上限
	begin time.Time     //週期開始時間
	cycle time.Duration //週期
	count int           //週期內請求次數
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

	if l.count == l.rate {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			l.reset(now)
			return true
		} else {
			return false
		}
	} else {
		l.count++
		return true
	}
}

func (l *LimitRate) reset(t time.Time) {
	l.begin = t
	l.count = 0
}

func (l *LimitRate) GetCount() int {
	return l.count
}
