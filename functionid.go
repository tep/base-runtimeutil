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
