package main

import (
	"gitlab.com/nibbleshift/argenv"
	"github.com/davecgh/go-spew/spew"
)

type MySettings struct {
	EthernetDevice string
	IpAddress string
	PortNumber int
	Username string
	Shell string
}

var settings *MySettings

func main() {
	argEnv := &argenv.ArgEnv{}
	settings = &MySettings{}

	argEnv.Load(settings)

	spew.Dump(settings)
}
