ArgEnv
===

ArgEnv is a simple package for quickly loading command line parameters, environment variables,
and default settings in your app.

The goal is to simplify loading configurable settings in your application.

## Usage Documentation

## Installation

Install Go. [Installation instructions here](http://golang.org/doc/install.html).


### Get the package

```
$ go get github.com/nibbleshift/argenv
```

```go
import (
  "github.com/nibbleshif/argenv" // imports as package "argenv"
)
...
```


### Example
```go
package main

import (
	"github.com/davecgh/go-spew/spew"
	"gitlab.com/nibbleshift/argenv"
)

type MySettings struct {
	EthernetDevice string `default: "eth0" description:"Specify NIC to configure"`
	IpAddress      string `default: "127.0.0.1" description:"IP Address to listen on"`
	PortNumber     int    `default: "80" description:"IP Address to listen on"`
	Username       string `default: "root" description:"Default user"`
	Shell          string `default: "/bin/bash" description:"Default Shell"`
}

var settings *MySettings

func main() {
	argEnv := &argenv.ArgEnv{}
	settings = &MySettings{}

	argEnv.Load(settings)
	spew.Dump(settings)
}
```
#### Running Example:
```
go run main.go -ip-address=192.168.100.1 -port-number=8080 \
	-username=steven -shell=/bin/bash -ethernet-device=eth1
```

#### Output 
```
(*main.MySettings)(0xc000074190)({
 EthernetDevice: (string) (len=4) "eth1",
 IpAddress: (string) (len=13) "192.168.100.1",
 PortNumber: (int) 8080,
 Username: (string) (len=6) "steven",
 Shell: (string) (len=9) "/bin/bash"
})
```
