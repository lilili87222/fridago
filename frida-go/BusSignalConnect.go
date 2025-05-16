package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"sync"
)

type BusOnMessageEventFunc func(sjson jsoniter.Any, data []byte)
type BusOnDetachEventFunc func()
type BusSignalConnect struct {
	onMessageSigs sync.Map
	onDetachSigs sync.Map
	rawBusPtr uintptr
}

func (s *BusSignalConnect) free() {
	if s.rawBusPtr!=0{
		//fmt.Println("BusSignalConnect gc")
		s.onMessageSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawBusPtr, key.(int64))
			bus_onMessageCallbackTable.Delete(key.(int64))
			return true
		})
		s.onDetachSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawBusPtr, key.(int64))
			bus_onDetachCallbackTable.Delete(key.(int64))
			return true
		})
		s.rawBusPtr=0
	}
}
func (s *BusSignalConnect) OnDetach(on BusOnDetachEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	bus_onDetachCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawBusPtr, "detached", bus_onDetachPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onDetachSigs.Store(sigid,userdata)
	return sigid
}
func (s *BusSignalConnect) OnMessage(on BusOnMessageEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	bus_onMessageCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawBusPtr, "message", bus_onMessagePtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onMessageSigs.Store(sigid,userdata)
	return sigid
}

var bus_onDetachCallbackTable=sync.Map{}
var bus_onMessageCallbackTable=sync.Map{}


func NewBusSignalConnect(rawPtr uintptr) *BusSignalConnect {
	sig := new(BusSignalConnect)
	sig.onMessageSigs = sync.Map{}
	sig.onDetachSigs = sync.Map{}
	sig.rawBusPtr = rawPtr
	return sig
}
