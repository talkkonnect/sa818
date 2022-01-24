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
	// Printing Variables to Screen

	log.Println("info: DMOSetup Values ", DMOSetup)

	//Sample Commands and Expected Results for Calling sa818

	err = sa818.Callsa818("InitComm", "(DMOCONNECT:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	}

	err = sa818.Callsa818("CheckVersion", "(VERSION:)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	}

	err = sa818.Callsa818("CheckRSSI", "(RSSI)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	}

	err = sa818.Callsa818("SetVolume", "(DMOSETVOLUME:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	}

	err = sa818.Callsa818("DMOSetupFilter", "(DMOSETFILTER:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
	}

	err = sa818.Callsa818("DMOSetupGroup", "(DMOSETGROUP:0)", DMOSetup)
	if err != nil {
		log.Println("error: From Module ", err)
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


