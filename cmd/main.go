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

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var DMOSetup sa818.DMOSetupStruct
	var initComm, checkVersion, checkRSSI, setFrequency, setFilter, setVolume bool

	DMOSetup.Band = 0
	DMOSetup.Rxfreq = 168.7750
	DMOSetup.Txfreq = 168.7750
	DMOSetup.Ctsstone = 0
	DMOSetup.Squelch = 1
	DMOSetup.Dcstone = 0
	DMOSetup.Predeemph = 0
	DMOSetup.Highpass = 0
	DMOSetup.Lowpass = 0
	DMOSetup.Volume = 8
	DMOSetup.SerialOptions.PortName = "/dev/ttyAMA0"
	DMOSetup.SerialOptions.BaudRate = 9600
	DMOSetup.SerialOptions.DataBits = 8
	DMOSetup.SerialOptions.StopBits = 1
	DMOSetup.SerialOptions.MinimumReadSize = 2
	DMOSetup.SerialOptions.InterCharacterTimeout = 200

	//enable the function you want to call to the module
	initComm = true
	checkVersion = true
	checkRSSI = true
	setFrequency = true
	setFilter = true
	setVolume = true

	if initComm {
		err := sa818.Callsa818("InitComm", DMOSetup)
		if err != nil {
			log.Println("info: SAModule Init Comm Error ", err)
		} else {
			log.Println("info: SAModule Init Comm OK ")
		}
	}

	if checkVersion {
		err := sa818.Callsa818("CheckVersion", DMOSetup)
		log.Println("info: CheckVersion ", err)
	}

	if checkRSSI {
		err := sa818.Callsa818("CheckRSSI", DMOSetup)
		log.Println("info: Check RSSI ", err)
	}

	if setFrequency {
		err := sa818.Callsa818("DMOSetupGroup", DMOSetup)
		if err != nil {
			log.Println("info: SAModule Set Frequecy Error ", err)
		} else {
			log.Println("info: SAModule Set Frequecy OK ")
		}
	}

	if setFilter {
		err := sa818.Callsa818("DMOSetupFilter", DMOSetup)
		if err != nil {
			log.Println("info: SAModule Setup Filter Error ", err)
		} else {
			log.Println("info: SAModule Setup Filter OK ")
		}
	}

	if setVolume {
		err := sa818.Callsa818("SetVolume", DMOSetup)
		if err != nil {
			log.Println("info: SAModule Set Volume Error ", err)
		} else {
			log.Println("info: SAModule Set Volume OK ")
		}
	}
}
