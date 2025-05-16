package main

import (
	"fmt"
	"freego/gofrida"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

func main() {
	//listDevices()
	listUsbDevices()
}
func listDevices() {

	mgr := gofrida.DeviceManager_Create()
	defer mgr.Close()
	ds, err := mgr.EnumerateDevices()
	if err != nil {
		panic(err)
	}
	for _, d := range ds {
		dtype := d.Type()
		tp := "未知"
		if dtype == 0 {
			tp = "本地"
		} else if dtype == 1 {
			tp = "远程"
		} else if dtype == 2 {
			tp = "usb"
		}
		fmt.Println(fmt.Sprintf("设备id:%s ,名称: %s 类型:%s", d.Id(), d.Name(), tp))
	}
}
func listUsbDevices() {
	dm := gofrida.DeviceManager_Create()
	//d, err := dm.FindDeviceByType(gofrida.DeviceType_USB, 1000)
	d, err := dm.FindDeviceById("8d8421ec89d1b323186308907a9cb50a1ea45cd8", 1000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {

		fmt.Println("Device found:", d.Name(), d.Description(), d.Id())
		//fmt.Println(d.QuerySystemParameters())
		//listApps(d)
		//app, err := findApps(d, "so.phonegame.hotgame")
		runJs(d, "Temu")

	}
}
func runJs(device *gofrida.Device, name string) {
	p, e := device.GetProcessByName(name, gofrida.ProcessMatchOptions{})
	if e != nil {
		fmt.Println("Error:", e)
		return
	}
	session, err := device.Attach(p.Pid(), gofrida.SessionOptions{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer session.Detach()
	sc, err := session.CreateScript(`
	console.log("ok111111111")
	`, gofrida.ScriptOptions{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sc.OnMessage(func(sjson jsoniter.Any, data []byte) {
		fmt.Println(sjson.ToString())
	})
	err = sc.Load()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sc.UnLoad()
}
func listApps(d *gofrida.Device) {
	apps, err := d.EnumerateApplications(gofrida.ApplicationQueryOptions{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, app := range apps {
		pid := app.Pid()
		if pid != 0 {
			fmt.Printf("名称:%-15s identifier:%-30s pid:%d\n", app.Name(), app.Identifier(), app.Pid())
		} else {
			fmt.Printf("名称:%-15s identifier:%-30s\n", app.Name(), app.Identifier())
		}
	}
}
func findApps(d *gofrida.Device, bundleId string) (*gofrida.ApplicationDetails, error) {
	bundleId = strings.TrimSpace(bundleId)
	apps, err := d.EnumerateApplications(gofrida.ApplicationQueryOptions{})
	if err != nil {
		return nil, err
	}
	for _, app := range apps {
		if app.Identifier() == bundleId {
			return app, nil
		}
	}
	return nil, fmt.Errorf("app not found with id :%s", bundleId)
}
