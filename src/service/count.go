package service

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
)

const (
	StatusSuccess         = 0
	StatusInvalidOffset   = 1
	StatusUnexpectedError = 2
)

func NewCountService() *ble.Service {
	n := 0
	s := ble.NewService(ble.MustParse("09fc95c0-c111-11e3-9904-0002a5d5c51b"))
	c := s.NewCharacteristic(ble.MustParse("11fac9e0-c111-11e3-9246-0002a5d5c51b"))
	c.HandleRead(
		ble.ReadHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
			fmt.Fprintf(rsp, "count: %d", n)
			n++
		}))

	s.NewCharacteristic(ble.MustParse("16fe0d80-c111-11e3-b8c8-0002a5d5c51b")).HandleWrite(
		ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
			log.Println("Wrote:", string(req.Data()))
		}))

	s.NewCharacteristic(ble.MustParse("1c927b50-c116-11e3-8a33-0800200c9a66")).HandleNotify(ble.NotifyHandlerFunc(
		func(r ble.Request, n ble.Notifier) {
			cnt := 0
			for {
				select {
				case <-n.Context().Done():
					fmt.Fprintf(n, "Count: %d", cnt)
					cnt++
					time.Sleep(time.Second)
				}
			}
		}))

	return s
}
