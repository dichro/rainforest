package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/dichro/rainforest"
	"github.com/golang/glog"
)

const tmpl = `Timestamp: {{ .TimeStamp }}
Device: {{ .DeviceMacId }}
Meter: {{ .MeterMacId }}
Demand: {{ .DemandKW }} kW
`

func main() {
	t := template.Must(template.New("").Parse(tmpl))
	flag.Parse()
	r := xml.NewDecoder(os.Stdin)
	for {
		m, err := rainforest.ParseEMU2Message(r)
		switch err {
		case nil:
		case io.EOF:
			return
		default:
			glog.Error(err)
		}
		switch msg := m.(type) {
		case *rainforest.InstantaneousDemand:
			t.Execute(os.Stdout, msg)
		default:
			fmt.Printf("unknown message: %#v\n", msg)
		}
	}
}
