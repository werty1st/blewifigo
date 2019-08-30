package service

import (
	"github.com/Wifx/gonetworkmanager/NetworkManager"
)

//https://github.com/BellerophonMobile/gonetworkmanager
func WpaService() {

	var devices []Device
	devices = NetworkManager.GetDevices()

}
