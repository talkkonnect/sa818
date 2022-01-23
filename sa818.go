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
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

type DMOSetupGroupStruct struct {
	Band     int
	Rxfreq   float32
	Txfreq   float32
	Ctsstone int
	Squelch  int
	Dcstone  int
}

type DMOSetupFilterStruct struct {
	Predeemph int
	Highpass  int
	Lowpass   int
}

var SerialOptions = serial.OpenOptions{
	PortName:        "/dev/ttyAMA0",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var DMOSetupGroup DMOSetupGroupStruct
var DMOSetupFilter DMOSetupFilterStruct

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

func SetFilter(DMOSetupFilterCommand DMOSetupFilterStruct) string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()
	serialWrite(fmt.Sprintf("AT+SETFILTER=%d,%d,%d", DMOSetupFilterCommand.Predeemph, DMOSetupFilterCommand.Highpass, DMOSetupFilterCommand.Lowpass)+"\r\n", port)
	time.Sleep(1 * time.Second)

	return serialRead(port)
}

func SetFrequency(DMOSetupGroupCommand DMOSetupGroupStruct) string {
	port, err := serial.Open(SerialOptions)
	if err != nil {
		log.Printf("error: Serial Open Failed with error %v", err)
		return "error"
	}

	defer port.Close()

	serialWrite(fmt.Sprintf("AT+DMOSETGROUP=%d,%.4f,%.4f,%04d,%d,%04d", DMOSetupGroupCommand.Band, DMOSetupGroupCommand.Rxfreq, DMOSetupGroupCommand.Txfreq, DMOSetupGroupCommand.Ctsstone, DMOSetupGroupCommand.Squelch, DMOSetupGroupCommand.Dcstone)+"\r\n", port)
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
	log.Printf("info: Wrote %v total of %v bytes.\n", stripMessage, n)
}

func serialRead(port io.ReadCloser) string {
	buf := make([]byte, 1048)
	for {
		n, err := port.Read(buf)
		if err != nil {
			return "error reading serial"
		} else {
			defer port.Close()
		}
		return string(buf[:n])
	}
}
