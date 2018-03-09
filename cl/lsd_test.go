package cl

import (
	"fmt"
)

// list all platforms and devices
func ExampleListDevices(){
	platforms, _ := GetPlatforms()
	for i, p := range platforms {
		fmt.Println("Platform", i, p.Name(), p.Version())
	}
	devices, _ := GetAllDevices(DeviceTypeAll)
	for i, d := range devices {
		fmt.Println("Device", i, d.Name(), d.DriverVersion())
	}
}
