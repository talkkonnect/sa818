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
 * sa818.go -> sa818 walkie talkie module golang library
 */

package sa818

import (
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

type DMOSetupStruct struct {
	Band          int
	Rxfreq        float32
	Txfreq        float32
	Ctsstone      int
	Squelch       int
	Dcstone       int
	Predeemph     int
	Highpass      int
	Lowpass       int
	Volume        int
	SerialOptions serial.OpenOptions
}

func Callsa818(sendCommand string, TDMOSetup DMOSetupStruct) error {
	var ExpectedAnswer string

	Port, err := serial.Open(TDMOSetup.SerialOptions)
	if err != nil {
		return fmt.Errorf("serial open failed with error %v", err)
	}

	defer Port.Close()

	switch sendCommand {
	case "InitComm":
		serialWrite("AT+DMOCONNECT\r\n", Port)
		ExpectedAnswer = "(DMOCONNECT:0)"
	case "CheckVersion":
		serialWrite("AT+VERSION\r\n", Port)
		ExpectedAnswer = "(VERSION:)"
	case "DMOSetupGroup":
		serialWrite(fmt.Sprintf("AT+DMOSETGROUP=%d,%.4f,%.4f,%04d,%d,%04d", TDMOSetup.Band, TDMOSetup.Rxfreq, TDMOSetup.Txfreq, TDMOSetup.Ctsstone, TDMOSetup.Squelch, TDMOSetup.Dcstone)+"\r\n", Port)
		ExpectedAnswer = "(DMOSETGROUP:0)"
	case "DMOSetupFilter":
		serialWrite(fmt.Sprintf("AT+SETFILTER=%d,%d,%d", TDMOSetup.Predeemph, TDMOSetup.Highpass, TDMOSetup.Lowpass)+"\r\n", Port)
		ExpectedAnswer = "(DMOSETFILTER:0)"
	case "SetVolume":
		serialWrite(fmt.Sprintf("AT+DMOSETVOLUME=%d\r\n", TDMOSetup.Volume), Port)
		ExpectedAnswer = "(DMOSETVOLUME:0)"
	case "CheckRSSI":
		serialWrite("RSSI?\r\n", Port)
		ExpectedAnswer = "(RSSI)"
	default:
		ExpectedAnswer = ""
		return errors.New("invalid command")
	}

	time.Sleep(1000 * time.Millisecond)
	AnswerGot := serialRead(Port)
	re := regexp.MustCompile(ExpectedAnswer)
	matched := re.MatchString(AnswerGot)
	if matched {
		if sendCommand == "CheckVersion" || sendCommand == "CheckRSSI" {
			return errors.New(AnswerGot)
		}
		return nil
	} else {
		return errors.New(AnswerGot)
	}
}

func serialWrite(message string, port io.WriteCloser) {
	b := []byte(message)
	n, err := port.Write(b)
	if err != nil {
		log.Println("error: cannot write")
		return
	}
	stripMessage := strings.TrimRight(message, "\r\n")
	if len(stripMessage) > 0 {
		log.Printf("debug: Wrote %v total of %v bytes.\n", stripMessage, n)
	}
}

func serialRead(port io.ReadCloser) string {
	buf := make([]byte, 1048)
	var n int = 0
	var err error
	for {
		n, err = port.Read(buf)
		if err != nil {
			return "error reading serial"
		} else {
			defer port.Close()
			return string(buf[:n])
		}
	}
}
