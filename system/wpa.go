package system

import (
	"errors"
	"log"

	nm "github.com/Wifx/gonetworkmanager"
)

func getssidsFromDevice(device nm.DeviceWireless) []string {

	log.Print("Scan Networks ...")
	aps, _ := device.GetAccessPoints()

	var ssids []string

	for index := 0; index < len(aps); index++ {
		ap := aps[index]
		ssid, _ := ap.GetPropertySSID()
		ssids = append(ssids, ssid)
		log.Printf("Found SSID: %s", ssid)
	}

	return ssids
}

//https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go
func GetSSIDs() ([]string, error) {

	var ssids []string
	//var mynm nm.NetworkManager
	//var error error
	mynm, error := nm.NewNetworkManager()
	if error != nil {
		log.Printf("Error loading NetworkManager %s.", error)
		return nil, error
	}

	devices, error := mynm.GetDevices()
	var wifiDevice nm.DeviceWireless
	for index := 0; index < len(devices); index++ {
		device := devices[index]

		_, wireless := device.(nm.DeviceWireless)
		if wireless {
			wifiDevice = device.(nm.DeviceWireless)
			//needs to scan with all devices
			myssids := getssidsFromDevice(wifiDevice)
			ssids = append(ssids, myssids...)
		}
	}

	if wifiDevice == nil {
		log.Printf("No Wifi device found.")
		return nil, errors.New("no wifi device found")
	}

	//todo remove duplicates
	// Request the device to scan. To know when the scan is finished, use the
	// "PropertiesChanged" signal from "org.freedesktop.DBus.Properties" to listen
	// to changes to the "LastScan" property.
	// https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go nm.Subscribe
	//assume one wifi card only

	return ssids, nil

}
