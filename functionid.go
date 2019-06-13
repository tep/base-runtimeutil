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

package runtimeutil

import (
	"fmt"
	"reflect"
	"runtime"
)

type FunctionInfo struct {
	name string
	file string
	line int
}

func (i *FunctionInfo) String() string {
	if i == nil {
		return "<unidentified>"
	}

	return fmt.Sprintf("%s [%s:%d]", i.name, i.file, i.line)
}

func (i *FunctionInfo) Name() string {
	if i == nil {
		return "<unknown>"
	}

	return i.name
}

func FuncID(fp interface{}) *FunctionInfo {
	p := reflect.ValueOf(fp).Pointer()
	if p == uintptr(0) {
		return nil
	}

	f := runtime.FuncForPC(p)
	if f == nil {
		return nil
	}

	id := &FunctionInfo{
		name: f.Name(),
	}

	id.file, id.line = f.FileLine(p)

	return id
}
