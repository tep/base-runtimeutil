// Copyright © 2018 Timothy E. Peoples <eng@toolman.org>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package runtimeutil // import "toolman.org/base/runtimeutil"

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
