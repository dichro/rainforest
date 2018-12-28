package main

import (
	"encoding/xml"
	"flag"
	"io"
	"os"

	"github.com/dichro/rainforest"
	"github.com/golang/glog"

	client "github.com/influxdata/influxdb/client/v2"
)

var (
	influxdb = flag.String("influxdb_addr", "", "HTTP address of influxdb")
	input    = flag.String("input", "", "serial device or file containing Rainforest XML")
	database = flag.String("influx_database", "rainforest", "influx database to write points to")
	series   = flag.String("influx_series", "power", "name for influx timeseries")
)

func main() {
	flag.Parse()
	f, err := os.Open(*input)
	if err != nil {
		glog.Exit(err)
	}
	r := xml.NewDecoder(f)

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: *influxdb,
	})
	if err != nil {
		glog.Exit(err)
	}
	defer c.Close()

	cfg := client.BatchPointsConfig{
		Database:  *database,
		Precision: "s",
	}
	glog.Info("starting loop")
	for {
		m, err := rainforest.ParseEMU2Message(r)
		switch err {
		case nil:
		case io.EOF:
			return
		default:
			glog.Error(err)
			continue
		}
		msg, ok := m.(*rainforest.InstantaneousDemand)
		if !ok {
			glog.Infof("ignoring unknown message %#v", m)
			continue
		}
		tags := map[string]string{
			"device": msg.DeviceMacId,
			"meter":  msg.MeterMacId,
		}
		fields := map[string]interface{}{"demand": msg.DemandKW()}
		pt, err := client.NewPoint(*series, tags, fields, msg.TimeStamp.Time)
		if err != nil {
			glog.Error(err)
			continue
		}
		bp, err := client.NewBatchPoints(cfg)
		if err != nil {
			glog.Exit(err)
		}
		bp.AddPoint(pt)
		if err := c.Write(bp); err != nil {
			glog.Error(err)
		}
		glog.Infof("wrote point %#v", msg)
	}
}
