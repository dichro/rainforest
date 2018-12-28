package rainforest

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// HexInt64 is an int64 that hex-decodes itself on UnmarshalXML.
type HexInt64 int64

func (n *HexInt64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	token, err := d.Token()
	if err != nil {
		return err
	}
	cd, ok := token.(xml.CharData)
	if !ok {
		return errors.New("expected CDATA")
	}
	// we don't own cd, so let's parse it before we call Skip or anything else in Decoder
	num, err := strconv.ParseInt(string(cd), 0, 64)
	if err != nil {
		return err
	}
	// are we guaranteed to have consumed all the CDATA now?
	if err := d.Skip(); err != nil {
		return err
	}
	*n = HexInt64(num)
	return err
}

type HexTime struct {
	time.Time
}

var hexTimeReference = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

func (t *HexTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var n HexInt64
	if err := n.UnmarshalXML(d, start); err != nil {
		return err
	}
	t.Time = hexTimeReference.Add(time.Duration(n) * time.Second)
	return nil
}

type InstantaneousDemand struct {
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           HexTime
	Demand              HexInt64
	Multiplier          HexInt64
	Divisor             HexInt64
	DigitsRight         HexInt64
	DigitsLeft          HexInt64
	SuppressLeadingZero string
}

// DemandKW returns the Demand expressed in kilowatts.
func (i InstantaneousDemand) DemandKW() float64 {
	mul := float64(i.Multiplier)
	if mul == 0 {
		mul = 1
	}
	div := float64(i.Divisor)
	if div == 0 {
		div = 1
	}
	return float64(i.Demand) * mul / div
}

type CurrentSummationDelivered struct {
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           HexTime
	SummationDelivered  HexInt64
	SummationReceived   HexInt64
	Multiplier          HexInt64
	Divisor             HexInt64
	DigitsRight         HexInt64
	DigitsLeft          HexInt64
	SuppressLeadingZero string
}

type TimeCluster struct {
	DeviceMacId string
	MeterMacId  string
	UTCTime     HexTime
	LocalTime   HexTime
}

// ParseEMU2Message returns the next EMU-2 message from this XML decoder.
func ParseEMU2Message(d *xml.Decoder) (interface{}, error) {
	var ret interface{}
	var start xml.StartElement
	for {
		t, err := d.Token()
		if err != nil {
			return ret, err
		}
		if s, ok := t.(xml.StartElement); ok {
			start = s
			break
		}
	}
	switch start.Name.Local {
	case "InstantaneousDemand":
		ret = &InstantaneousDemand{}
	case "CurrentSummationDelivered":
		ret = &CurrentSummationDelivered{}
	case "TimeCluster":
		ret = &TimeCluster{}
	default:
		err := fmt.Errorf("unknown message type %q", start.Name)
		d.Skip()
		return ret, err
	}
	return ret, d.DecodeElement(ret, &start)
}
