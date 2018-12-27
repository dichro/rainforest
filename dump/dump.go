package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dichro/rainforest"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	r := xml.NewDecoder(os.Stdin)
	var m rainforest.EMU2Message
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
