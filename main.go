package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-ble/ble"

	//"github.com/werty1st/blewifigo/system"
	"github.com/werty1st/blewifigo/dev"
	"github.com/werty1st/blewifigo/service"

	system "github.com/werty1st/blewifigo/mock"
)

var (
	device = flag.String("device", "default", "implementation of ble")
	du     = flag.Duration("du", 5*time.Second, "advertising duration, 0 for indefinitely")
)

func main() {
	d, err := dev.NewDevice("default")
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}
	ble.SetDefaultDevice(d)

	ssids, _ := system.GetSSIDs()

	log.Print("ssids:", len(ssids))

	d.AddService(service.NewGattService())
	d.AddService(service.NewCountService())
	d.AddService(service.NewBatteryService())

	// Register optional handlers.

	// d.Handle(
	// 	gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
	// 	gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	// )

	// // A mandatory handler for monitoring device state.
	// onStateChanged := func(d gatt.Device, s gatt.State) {
	// 	fmt.Printf("State: %s\n", s)
	// 	switch s {
	// 	case gatt.StatePoweredOn:
	// 		// Setup GAP and GATT services for Linux implementation.
	// 		// OS X doesn't export the access of these services.
	// 		// d.AddService(service.NewGapService("Gopher")) // no effect on OS X
	// 		d.AddService(service.NewGattService()) // no effect on OS X

	// 		// A simple count service for demo.
	// 		s1 := service.NewCountService()
	// 		d.AddService(s1)

	// 		// A fake battery service for demo.
	// 		s2 := service.NewBatteryService()
	// 		d.AddService(s2)

	// 		// Advertise device name and service's UUIDs.
	// 		d.AdvertiseNameAndServices("Gopher", []gatt.UUID{s1.UUID(), s2.UUID()})

	// 		// Advertise as an OpenBeacon iBeacon
	// 		d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)

	// 	default:
	// 	}
	// }

	// d.Init(onStateChanged)

	//run loop?
	select {}
}
