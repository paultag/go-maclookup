package maclookup

import (
	"bytes"
	"fmt"
	"net"
)

var NotFound error = fmt.Errorf("No matching vendor found")

type Vendor struct {
	Prefix []byte
	Name   string
}

func (v Vendor) String() string {
	if v.Prefix == nil {
		return fmt.Sprintf("Prefix: n/a,    Name: n/a")
	}
	return fmt.Sprintf("Prefix: %x, Name: %s", v.Prefix, v.Name)
}

func (v Vendor) Contains(mac net.HardwareAddr) bool {
	return bytes.HasPrefix(mac, v.Prefix)
}

type Vendors []Vendor

func Lookup(mac net.HardwareAddr) (*Vendor, error) {
	if mac == nil {
		return &Vendor{Name: ""}, nil
	}
	for _, vendor := range List {
		if vendor.Contains(mac) {
			return &vendor, nil
		}
	}
	return nil, NotFound
}
