package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
)

type InstantaneousDemand struct {
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           string
	Demand              string
	Multiplier          string
	Divisor             string
	DigitsRight         string
	DigitsLeft          string
	SuppressLeadingZero string
}

type EMU2Message struct {
	InstantaneousDemand
}

func main() {
	flag.Parse()
	/*
		f, err := os.Open("/dev/ttyACM0")
		if err != nil {
			glog.Exit(err)
		}
	*/
	f := os.Stdin
	r := xml.NewDecoder(f)
	var m EMU2Message
	for {
		switch err := r.Decode(&m); err {
		case io.EOF:
			return
		case nil:
			fmt.Printf("%#v\n", m)
		default:
			glog.Error(err)
		}
	}
}
