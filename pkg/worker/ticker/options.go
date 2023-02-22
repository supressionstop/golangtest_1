package ticker

import "time"

type Option func(*Worker)

func Interval(interval time.Duration) Option {
	return func(worker *Worker) {
		worker.interval = interval
	}
}
