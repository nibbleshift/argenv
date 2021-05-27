package main

import (
	"github.com/davecgh/go-spew/spew"
	"gitlab.com/nibbleshift/argenv"
)

type MySettings struct {
	EthernetDevice string `default:"eth0" description:"Specify NIC to configure"`
	IpAddress      string `default:"127.0.0.1" description:"IP Address to listen on"`
	PortNumber     int    `default:"80" description:"IP Address to listen on"`
	Username       string `default:"root" description:"Default user"`
	Shell          string `default:"/bin/bash" description:"Default Shell"`
}

var settings *MySettings

func main() {
	argEnv := &argenv.ArgEnv{}
	settings = &MySettings{}

	argEnv.Load(settings)
	spew.Dump(settings)
}
