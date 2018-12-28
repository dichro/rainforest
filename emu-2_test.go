package rainforest

import (
	"encoding/xml"
	"testing"
)

const instantaneousDemand = `<InstantaneousDemand>
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

func TestInstantaneousDemand(t *testing.T) {
	var m InstantaneousDemand
	if err := xml.Unmarshal([]byte(instantaneousDemand), &m); err != nil {
		t.Fatalf("unmarshal failed: %s", err)
	}
}

const currentSummationDelivered = `<CurrentSummationDelivered>
<DeviceMacId>0xd8d5b9000000f919</DeviceMacId>
<MeterMacId>0x00135003004653a5</MeterMacId>
<TimeStamp>0x23b81fe2</TimeStamp>
<SummationDelivered>0x0000000000c96379</SummationDelivered>
<SummationReceived>0x0000000000000000</SummationReceived>
<Multiplier>0x00000001</Multiplier>
<Divisor>0x000003e8</Divisor>
<DigitsRight>0x03</DigitsRight>
<DigitsLeft>0x0f</DigitsLeft>
<SuppressLeadingZero>Y</SuppressLeadingZero>
</CurrentSummationDelivered>`

func TestCurrentSummationDelivered(t *testing.T) {
	var m CurrentSummationDelivered
	if err := xml.Unmarshal([]byte(currentSummationDelivered), &m); err != nil {
		t.Fatalf("unmarshal failed: %s", err)
	}
}

const timeCluster = `<TimeCluster>
<DeviceMacId>0xd8d5b9000000f919</DeviceMacId>
<MeterMacId>0x00135003004653a5</MeterMacId>
<UTCTime>0x23b81f68</UTCTime>
<LocalTime>0x23b7aee8</LocalTime>
</TimeCluster>`

func TestTimeCluster(t *testing.T) {
	var m TimeCluster
	if err := xml.Unmarshal([]byte(timeCluster), &m); err != nil {
		t.Fatalf("unmarshal TimeCluster failed: %s", err)
	}
}
