package pvz_service

import "time"

type TimerInterface interface {
	Now() time.Time
}

func NewTimer() TimerInterface {
	return &Timer{}
}

type Timer struct{}

func (*Timer) Now() time.Time {
	return time.Now()
}
