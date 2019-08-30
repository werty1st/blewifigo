package system

import (
	"errors"
	"log"

	. "github.com/Wifx/gonetworkmanager"
)

//https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go
func GetSSIDs() (DeviceWireless, error) {

	//var mynm nm.NetworkManager
	//var error error
	mynm, error := NewNetworkManager()
	if error != nil {
		log.Printf("Error loading NetworkManager %s.", error)
		return nil, error
	}

	devices, error := mynm.GetDevices()
	var wifiDevice DeviceWireless
	for index := 0; index < len(devices); index++ {
		device := devices[index]

		_, wireless := device.(DeviceWireless)
		if wireless {
			wifiDevice = device.(DeviceWireless)
			break
		}
	}

	if wifiDevice == nil {
		log.Printf("No Wifi device found.")
		return nil, errors.New("No Wifi device found.")
	}

	log.Printf("Scan Networks ...")
	aps, error := wifiDevice.GetAccessPoints()

	for index := 0; index < len(aps); index++ {
		ap := aps[index]
		ssid, _ := ap.GetPropertySSID()
		log.Printf(ssid)
	}

	return wifiDevice, nil

}
