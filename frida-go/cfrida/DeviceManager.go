package cfrida

import "errors"

func Frida_device_manager_new() uintptr {
	r, _, _ := frida_device_manager_new.Call()
	return r
}
func Frida_remote_device_options_new() uintptr {
	r, _, _ := frida_remote_device_options_new.Call()
	return r
}
func Frida_remote_device_options_set_certificate(obj uintptr,val uintptr) {
	frida_remote_device_options_set_certificate.Call(obj,val)
}
func Frida_remote_device_options_set_origin(obj uintptr,val string) {
	frida_remote_device_options_set_origin.Call(obj,GoStrToCStr(val))
}
func Frida_remote_device_options_set_token(obj uintptr,val string) {
	frida_remote_device_options_set_token.Call(obj,GoStrToCStr(val))
}
func Frida_remote_device_options_set_keepalive_interval(obj uintptr,n int) {
	frida_remote_device_options_set_keepalive_interval.Call(obj, uintptr(n))
}
func Frida_device_manager_close_sync(obj uintptr, cancellable uintptr)error{
	gerr:=MakeGError()
	frida_device_manager_close_sync.Call(obj,cancellable, gerr.Input())
	return gerr.ToError()
}
func Frida_device_manager_enumerate_devices_sync(obj uintptr, _cancellable uintptr) (uintptr,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_manager_enumerate_devices_sync.Call(obj, _cancellable, gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_list_size(obj uintptr) int {
	r, _, _ := frida_device_list_size.Call(obj)
	return int(r)
}
func Frida_device_list_get(obj uintptr,index int) uintptr {
	r, _, _ := frida_device_list_get.Call(obj,uintptr(index))
	return r
}
func Frida_device_manager_add_remote_device_sync(obj uintptr,address string,ops uintptr,cancellable uintptr) (uintptr,error){
	if address==""{
		return 0,errors.New("address is emtry")
	}
	gerr:=MakeGError()
	r,_,_:=frida_device_manager_add_remote_device_sync.Call(obj,GoStrToCStr(address),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_manager_remove_remote_device_sync(obj uintptr,address string,cancellable uintptr) error{
	if address==""{
		return errors.New("address is emtry")
	}
	gerr:=MakeGError()
	frida_device_manager_add_remote_device_sync.Call(obj,GoStrToCStr(address),cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_device_manager_get_device_by_id_sync(obj uintptr,id string,timeout int,cancellable uintptr) (uintptr,error){
	if id==""{
		return 0,errors.New("id is emtry")
	}
	gerr:=MakeGError()
	r,_,_:=frida_device_manager_get_device_by_id_sync.Call(obj,GoStrToCStr(id), uintptr(timeout),cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_manager_get_device_by_type_sync(obj uintptr,tp int,timeout int,cancellable uintptr) (uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_manager_get_device_by_type_sync.Call(obj, uintptr(tp), uintptr(timeout),cancellable,gerr.Input())
	return r,gerr.ToError()
}

func Frida_device_manager_find_device_by_id_sync(obj uintptr,id string,timeout int,cancellable uintptr) (uintptr,error){
	if id==""{
		return 0,errors.New("id is emtry")
	}
	gerr:=MakeGError()
	r,_,_:=frida_device_manager_find_device_by_id_sync.Call(obj,GoStrToCStr(id), uintptr(timeout),cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_manager_find_device_by_type_sync(obj uintptr,tp int,timeout int,cancellable uintptr) (uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_manager_find_device_by_type_sync.Call(obj, uintptr(tp), uintptr(timeout),cancellable,gerr.Input())
	return r,gerr.ToError()
}

