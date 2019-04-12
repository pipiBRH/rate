package rate

import (
	"testing"
	"time"
)

func Test_Allow_1(t *testing.T) {
	limiter := NewLimiter(5, time.Minute)
	for i := 0; i < 5; i++ {
		limiter.Allow()
	}
	if limiter.Allow() == true {
		t.Error("status error")
	}
}

func Test_Allow_2(t *testing.T) {
	limiter := NewLimiter(5, time.Minute)
	for i := 0; i < 4; i++ {
		limiter.Allow()
	}
	if limiter.Allow() == false {
		t.Error("status correct")
	}
}

func Test_GetCount(t *testing.T) {
	limiter := NewLimiter(10, time.Minute)
	for i := 0; i < 10; i++ {
		limiter.Allow()
	}
	if limiter.GetCount() != 10 {
		t.Error("status correct")
	}
}

func Test_Reste(t *testing.T) {
	limiter := NewLimiter(10, time.Minute)
	for i := 0; i < 10; i++ {
		limiter.Allow()
	}
	if limiter.GetCount() != 10 {
		t.Error("status correct")
	}
}
