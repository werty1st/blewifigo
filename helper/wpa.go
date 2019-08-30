package service

import (
	"log"

	nm "github.com/Wifx/gonetworkmanager"
)

//https://github.com/Wifx/gonetworkmanager/blob/master/NetworkManager.go
func WpaService() {

	//var mynm nm.NetworkManager
	//var error error
	mynm, error := nm.NewNetworkManager()
	if error != nil {
		log.Println("Error loading NetworkManager %s", error.Error)
		return
	}

	devices, error := mynm.GetAllDevices()

	log.Println("Loaded NetworkManager %i", len(devices))

}
