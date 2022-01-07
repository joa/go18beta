package future

import "errors"

var (
	ErrTimeout       = errors.New("timeout exceeded")
	ErrChannelClosed = errors.New("channel closed")
)
