package operations

import "time"

type ApiReader interface {
	GetUrl(d time.Time) (string)
	MaxConcurrentRequests() int
}