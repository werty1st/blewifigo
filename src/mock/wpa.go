package system

import (
	"log"

)

func getssidsFromDevice() []string {

	log.Print("Scan Networks ...")
	ssids := []string{ "SSID A", "SSID B", "SSID C" }

	return ssids
}

//https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go
func GetSSIDs() ([]string, error) {

	var ssids []string
	//var mynm nm.NetworkManager


	//needs to scan with all devices
	ssids = getssidsFromDevice()

	//todo remove duplicates
	// Request the device to scan. To know when the scan is finished, use the
	// "PropertiesChanged" signal from "org.freedesktop.DBus.Properties" to listen
	// to changes to the "LastScan" property.
	// https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go nm.Subscribe
	//assume one wifi card only

	return ssids, nil

}
