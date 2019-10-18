package service

import (
	"log"

	"github.com/go-ble/ble"
)

var (
	attrGATTUUID           = ble.UUID16(0x1801)
	attrServiceChangedUUID = ble.UUID16(0x2A05)
)

// NOTE: OS X provides GAP and GATT services, and they can't be customized.
// For Linux/Embedded, however, this is something we want to fully control.
func NewGattService() *ble.Service {
	s := ble.NewService(attrGATTUUID)
	s.NewCharacteristic(attrServiceChangedUUID).HandleNotify(ble.NotifyHandlerFunc(
		func(r ble.Request, n ble.Notifier) {
			go func() {
				log.Printf("TODO: indicate client when the services are changed")
			}()
		}))
	return s
}
