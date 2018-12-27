package rainforest

import (
	"encoding/xml"
	"errors"
	"strconv"
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

type InstantaneousDemand struct {
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           HexInt64
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

type EMU2Message struct {
	InstantaneousDemand
	// CurrentSummationDelivered
	// TimeCluster
}
