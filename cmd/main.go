/*
 * sa8118 golang library for controlling nicerf walkie talkie modules with go
 * Copyright (C) 2018-2022, Suvir Kumar <suvir@talkkonnect.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Software distributed under the License is distributed on an "AS IS" basis,
 * WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
 * for the specific language governing rights and limitations under the
 * License.
 *
 * talkkonnect is the based on talkiepi and barnard by Daniel Chote and Tim Cooper
 *
 * The Initial Developer of the Original Code is
 * Suvir Kumar <suvir@talkkonnect.com>
 * Portions created by the Initial Developer are Copyright (C) Suvir Kumar. All Rights Reserved.
 *
 * Contributor(s):
 *
 * Suvir Kumar <suvir@talkkonnect.com>
 *
 * My Blog is at www.talkkonnect.com
 * The source code is hosted at github.com/talkkonnect
 *
 * main.go -> main demo function in sa818 go library for controlling nicerf rf walkietalkie moduless
 */

package main

import (
	"log"
	"github.com/talkkonnect/sa818"
)

var DMOSetup sa818.DMOSetupStruct

func main() {

	DMOSetup.Band = 0
	DMOSetup.Rxfreq = 168.7750
	DMOSetup.Txfreq = 168.7750
	DMOSetup.Ctsstone = 0
	DMOSetup.Squelch = 0
	DMOSetup.Dcstone = 0
	DMOSetup.Predeemph = 0
	DMOSetup.Highpass = 0
	DMOSetup.Lowpass = 0
	DMOSetup.Volume = 4
	DMOSetup.PortName = "/dev/ttyAMA0"
	DMOSetup.BaudRate = 9600
	DMOSetup.DataBits = 8
	DMOSetup.StopBits = 1

	var err error
	var message string

	message, err = sa818.Callsa818("InitComm", "(DMOCONNECT:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

	message, err = sa818.Callsa818("CheckVersion", "(VERSION:)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

	message, err = sa818.Callsa818("CheckRSSI", "(RSSI)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

	message, err = sa818.Callsa818("SetVolume", "(DMOSETVOLUME:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

	message, err = sa818.Callsa818("DMOSetupFilter", "(DMOSETFILTER:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

	message, err = sa818.Callsa818("DMOSetupGroup", "(DMOSETGROUP:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	} else {
		log.Println("info: sa818 says ", message)
	}

}
