package main

import (
	"airship.go"
	"log"
)

const AppKey = "YOUR_APP_KEY"
const AppMasterSecret = "YOUR_APP_MASTER_SECRET"
const DeviceToken = "YOUR_DEVICE_TOKEN"

func main() {
	app := airship.App{Key: AppKey, MasterSecret: AppMasterSecret}ss
	aps_data := airship.APS{Alert: "hi!"}
	data := airship.PushData{APS: aps_data, DeviceTokens: []string{DeviceToken}}
	err := app.Push(data); if err != nil {
		log.Print(err)
	}
}