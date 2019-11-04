package service

import "github.com/go-ble/ble"

// NewBatteryService ...
func NewBatteryService() *ble.Service {
	lv := byte(100)
	s := ble.NewService(ble.UUID16(0x180F))
	c := s.NewCharacteristic(ble.UUID16(0x2A19))
	c.HandleRead(
		ble.ReadHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
			rsp.Write([]byte{lv})
			lv--
		}))

	// Characteristic User Description
	c.NewDescriptor(ble.UUID16(0x2901)).SetValue([]byte("Battery level between 0 and 100 percent"))

	// Characteristic Presentation Format
	c.NewDescriptor(ble.UUID16(0x2904)).SetValue([]byte{4, 1, 39, 173, 1, 0, 0})

	return s
}

// func NewBatteryService() *gatt.Service {
// 	lv := byte(100)
// 	s := gatt.NewService(gatt.UUID16(0x180F))
// 	c := s.AddCharacteristic(gatt.UUID16(0x2A19))
// 	c.HandleReadFunc(
// 		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
// 			rsp.Write([]byte{lv})
// 			lv--
// 		})

// 	// FIXME: this cause connection interrupted on Mac.
// 	// Characteristic User Description
// 	// c.AddDescriptor(gatt.UUID16(0x2901)).SetValue([]byte("Battery level between 0 and 100 percent"))

// 	// Characteristic Presentation Format
// 	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{4, 1, 39, 173, 1, 0, 0})

// 	return s
// }
