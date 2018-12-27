package main

import (
	"encoding/xml"
	"testing"
)

const emu2message = `<InstantaneousDemand>
  <DeviceMacId>0xd8d5b9000000f919</DeviceMacId>
  <MeterMacId>0x00135003004653a5</MeterMacId>
  <TimeStamp>0x23b810ea</TimeStamp>
  <Demand>0x0003c4</Demand>
  <Multiplier>0x00000001</Multiplier>
  <Divisor>0x000003e8</Divisor>
  <DigitsRight>0x03</DigitsRight>
  <DigitsLeft>0x0f</DigitsLeft>
  <SuppressLeadingZero>Y</SuppressLeadingZero>
</InstantaneousDemand>`

func TestEMU2Message(t *testing.T) {
	var m EMU2Message
	if err := xml.Unmarshal([]byte(emu2message), &m); err != nil {
		t.Fatalf("unmarshal failed: %s", err)
	}
}
