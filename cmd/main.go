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

var DMOSetupGroup sa818.DMOSetupGroupStruct
var DMOSetupFilter sa818.DMOSetupFilterStruct
var volume int = 5

func main() {

	DMOSetupGroup.Band = 0
	DMOSetupGroup.Rxfreq = 168.7750
	DMOSetupGroup.Txfreq = 168.7750
	DMOSetupGroup.Ctsstone = 0
	DMOSetupGroup.Squelch = 0
	DMOSetupGroup.Dcstone = 0

	DMOSetupFilter.Predeemph = 0
	DMOSetupFilter.Highpass = 0
	DMOSetupFilter.Lowpass = 0

	log.Println(DMOSetupGroup)
	log.Println(DMOSetupFilter)
	log.Println(volume)

	//log.Println(sa818.InitComm())
	//log.Println(sa818.CheckVersion())
	log.Println(sa818.SetFrequency(DMOSetupGroup))
	//log.Println(sa818.SetFilter(DMOSetupFilter))
	log.Println(sa818.SetVolume(volume))
	//log.Println(sa818.SetCloseTailTone(1))
	//log.Println(sa818.CheckRSSI())

}
