package main

import (
	"encoding/xml"
	"flag"
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
		var m rainforest.EMU2Message
		switch err := r.Decode(&m); err {
		case io.EOF:
			return
		case nil:
			if err := t.Execute(os.Stdout, m.InstantaneousDemand); err != nil {
				glog.Exit(err)
			}
		default:
			glog.Error(err)
		}
	}
}
