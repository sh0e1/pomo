package timer

import "time"

type Config struct {
	WorkInterval       time.Duration
	ShortBreakInterval time.Duration
	LongBreakInterval  time.Duration
	Rounds             int
}
