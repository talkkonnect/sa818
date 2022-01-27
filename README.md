upSA818 golang library for serial control
=======================================

This library written in [Go programming language](https://golang.org/) to control a rf sa818 walkie talkie module

![image](https://raw.github.com/talkkonnect/sa818/master/images/sa818.jpg)

Compatibility
-------------
Tested on Raspberry PI 3 (model B+)

Golang usage
------------

```go

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

```

Installation
------------

```bash
$ go get -u github.com/talkkonnect/sa818
```

Credits
-------

This project is inspired by sa818 python program available at https://pypi.org/project/sa818/

Contact
-------

Please use [Github issue tracker](https://github.com/talkkonnect/max7219/issues) for filing bugs or feature requests.

License
-------

sa818 is licensed under MIT License.


