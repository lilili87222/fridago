//go:build windows && !posix
// +build windows,!posix

package gofrida

import (
	"syscall"
)

var bus_onDetachPtr = syscall.NewCallbackCDecl(func(sc uintptr, userdata uintptr) uintptr {
	return bus_onDetachCallBack(sc, userdata)
})
var bus_onMessagePtr = syscall.NewCallbackCDecl(func(sc uintptr, rawjson uintptr, rawdata uintptr, userdata uintptr) uintptr {
	return bus_onMessageCallBack(sc, rawjson, rawdata, userdata)
})
