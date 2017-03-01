package keepalive

import (
	"time"
)

// ClientParameters is used to set keepalive parameters for the client side.
type ClientParameters struct {
	// After a duration of this time the client pings the server to see if the transport is still alive.
	Time time.Duration
	// After having pinged for keepalive check, the client waits for a duration of keepalive_timeout before closing the transport.
	Timeout time.Duration
	// If true, client runs keepalive checks even with no active RPCs.
	PermitWithoutStream bool
}
