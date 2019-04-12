package rate

import (
	"fmt"
	"net/http"
	"strings"
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
	var lr LimitRate
	lr.Set(r, cycle)
	return &lr
}

func (l *LimitRate) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count == l.rate {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			l.Reset(now)
			return true
		} else {
			return false
		}
	} else {
		l.count++
		return true
	}
}

func (l *LimitRate) Set(r int, cycle time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

func (l *LimitRate) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}

func (l *LimitRate) GetCount() int {
	return l.count
}
