package now

import (
	"time"
)

const (
	op = "service/now"
)

var (
	// Now returns time.Time
	Now func() time.Time = time.Now
)
