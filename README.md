SA818 golang library for serial control
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


