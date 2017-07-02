package runtimeutil  // import "toolman.org/base/runtimeutil"

import (
	"errors"
	"runtime"
)

func CalledFrom() (*runtime.Frame, error) {
	pcs := make([]uintptr, 1)
	if n := runtime.Callers(3, pcs); n == 0 {
		return nil, errors.New("no callers found")
	}

	frames := runtime.CallersFrames(pcs)
	if frames == nil {
		return nil, errors.New("no call frames")
	}

	cf, _ := frames.Next()

	return &cf, nil
}
