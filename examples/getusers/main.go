package main

import (
	"io/ioutil"
	"log"

	"github.com/jkl-sp/onvif"
	"github.com/jkl-sp/onvif/device"
)

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.149",
		Username: "administrator",
		Password: "Password1!",
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}
	// CreateUsers
	res, err := dev.CallMethod(device.GetUsers{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := ioutil.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
