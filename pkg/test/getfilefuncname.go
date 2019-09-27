package test

import (
	"path"
	"reflect"
	"runtime"
	"strings"
)

// GetFileFuncName takes a function and returns a string formatted as
// "basedir/filename/function" to make it easier to identify the
// source in printed test names (which will consequently appear as
// "TestFoo/basedir/filename/function")
func GetFileFuncName(f interface{}) string {
	pc := reflect.ValueOf(f).Pointer()
	fpc := runtime.FuncForPC(pc)
	parts := strings.Split(path.Base(fpc.Name()), ".")

	file, _ := fpc.FileLine(pc)
	base := path.Base(file)

	// parts[0] => basedir
	// base     => filename
	// parts[1] => function
	name := path.Join(parts[0], base, parts[1])
	return name
}
