// Copyright 2011 Urban Airship, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


// Package airship allows you to easily communicate with the Urban Airship
// ( http://urbanairship.com/ ).
package airship

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"io/ioutil"
	"json"
	"os"
)

var UAClient = &http.Client{}

// APS represents an iOS payload - alert, sound, badge.
type APS struct {
	Alert string `json:"alert,omitempty"`
	Sound string `json:"sound,omitempty"`
	Badge int `json:"badge,omitempty"`
}

type Android struct {
	Alert string `json:"alert,omitempty"`
}

// PushData is a struct that represents a payload.
type PushData struct {
	APS APS `json:"aps,omitempty"`
	Android Android `json:"android,omitempty"`
	DeviceTokens []string `json:"device_tokens,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
}

// App represents an Urban Airship application.
type App struct {
	Key string
	MasterSecret string
	ServerUrl string
}

func (app *App) deliverPayload(url string, payload io.Reader) os.Error {
	if (app.ServerUrl == "") {
		app.ServerUrl = "https://go.urbanairship.com"
	}
	apiEndpoint := app.ServerUrl + url
	req, err := http.NewRequest("POST", apiEndpoint, payload); if err != nil {
		return err
	}
	req.SetBasicAuth(app.Key, app.MasterSecret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := UAClient.Do(req); if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		respString, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return os.NewError(fmt.Sprintf("Hit a non-200 response from UA with a status code of %s: %s\n", resp.StatusCode, respString))
	}
	return nil
}

// Takes data, marshals it, and sends it along to the broadcast API endpoint.
func (app *App) Broadcast(data PushData) os.Error {
	json_data, err := json.Marshal(data); if err != nil {
		return err
	}
	payload := bytes.NewBuffer(json_data)
	return app.deliverPayload("/api/push/broadcast/", payload)
}

// Takes data, marshals it, and sends it along to the push API endpoint.
func (app *App) Push(data PushData) os.Error {
	json_data, err := json.Marshal(data); if err != nil {
		return err
	}
	payload := bytes.NewBuffer(json_data)
	return app.deliverPayload("/api/push/", payload)
}