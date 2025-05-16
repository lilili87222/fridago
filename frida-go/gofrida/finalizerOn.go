package gofrida

import "runtime"

func setFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, finalizer)
}
