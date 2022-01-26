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
	Band      int
	Rxfreq    float32
	Txfreq    float32
	Ctsstone  int
	Squelch   int
	Dcstone   int
	Predeemph int
	Highpass  int
	Lowpass   int
	Volume    int
	PortName  string
	BaudRate  uint
	DataBits  uint
	StopBits  uint
}

var DMOSetup DMOSetupStruct
var SA818Answer string

var SerialOptions serial.OpenOptions

func CheckRSSI() string {

	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite("RSSI?\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func SetCloseTailTone(tail int) string {
	if !(tail == 0 || tail == 1) {
		log.Printf("error: Valid Tail Tone Ranges (0-1) %v is out of range\n", tail)
		return "error"
	}
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite(fmt.Sprintf("AT+SETTAIL=%d\r\n", tail), port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func SetVolume(volume int) string {
	if volume < 0 || volume > 8 {
		log.Printf("error: Valid Volume Ranges (0-8) %v is out of range\n", volume)
		return "error"
	}
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite(fmt.Sprintf("AT+DMOSETVOLUME=%d\r\n", volume), port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func SetFilter(DMOSetupCommand DMOSetupStruct) string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()
	serialWrite(fmt.Sprintf("AT+SETFILTER=%d,%d,%d", DMOSetupCommand.Predeemph, DMOSetupCommand.Highpass, DMOSetupCommand.Lowpass)+"\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func SetFrequency(DMOSetupCommand DMOSetupStruct) string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite(fmt.Sprintf("AT+DMOSETGROUP=%d,%.4f,%.4f,%04d,%d,%04d", DMOSetupCommand.Band, DMOSetupCommand.Rxfreq, DMOSetupCommand.Txfreq, DMOSetupCommand.Ctsstone, DMOSetupCommand.Squelch, DMOSetupCommand.Dcstone)+"\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func CheckVersion() string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite("AT+VERSION\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)

}

func InitComm() string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite("AT+DMOCONNECT\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)

}

func serialWrite(message string, port io.WriteCloser) {
	b := []byte(message)
	n, err := port.Write(b)
	if err != nil {
		log.Println("error: cannot write")
		return
	}
	stripMessage := strings.TrimRight(message, "\r\n")
	log.Printf("debug: Wrote %v total of %v bytes.\n", stripMessage, n)
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

func Callsa818(sendCommand string, expectedAnswer string, DMOSetup DMOSetupStruct) (string, error) {

	SerialOptions.PortName = DMOSetup.PortName
	SerialOptions.BaudRate = DMOSetup.BaudRate
	SerialOptions.DataBits = DMOSetup.DataBits
	SerialOptions.DataBits = DMOSetup.DataBits
	SerialOptions.StopBits = DMOSetup.StopBits
	SerialOptions.MinimumReadSize = 4

	var ErrorMessage error
	var Message string

	switch sendCommand {
	case "InitComm":
		SA818Answer = InitComm()
		Message = "Successfully Initalized Module"
		ErrorMessage = errors.New("cannot initalize module")
	case "CheckVersion":
		SA818Answer = CheckVersion()
		Message = "Successfully Checked Version"
		ErrorMessage = errors.New("cannot check version")
	case "DMOSetupGroup":
		SA818Answer = SetFrequency(DMOSetup)
		Message = "DMOSetupGroup OK"
		ErrorMessage = errors.New("cannot setup frequency")
	case "DMOSetupFilter":
		SA818Answer = SetFilter(DMOSetup)
		Message = "DMOSetupFilter OK"
		ErrorMessage = errors.New("cannot setup filter")
	case "SetVolume":
		SA818Answer = SetVolume(DMOSetup.Volume)
		Message = "Set Volume OK"
		ErrorMessage = errors.New("cannot set volume")
	case "SetCloseTailTone":
		SA818Answer = SetCloseTailTone(1)
		Message = "Set CloseTailTone OK"
		ErrorMessage = errors.New("cannot close tail tone")
	case "SetOpenTailTone":
		SA818Answer = SetCloseTailTone(0)
		Message = "Set OpenTailTone OK"
		ErrorMessage = errors.New("cannot open tail tone")
	case "CheckRSSI":
		SA818Answer = CheckRSSI()
		Message = "Check RSSI OK"
		ErrorMessage = errors.New("cannot check rssi")
	default:
		return "Command Not OK", errors.New("command not defined")
	}

	re := regexp.MustCompile(expectedAnswer)
	matched := re.MatchString(SA818Answer)
	if matched {
//		log.Printf("debug: OK Response From sa818 %v\n", SA818Answer)
		time.Sleep(800 * time.Millisecond)
		return Message, nil
	} else {
//		log.Println("error: Fail Response From sa818 ", SA818Answer)
		return "Error", ErrorMessage
	}
}
