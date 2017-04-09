package maclookup

import (
	"bytes"
	"fmt"
	"net"
)

type Vendor struct {
	Prefix []byte
	Name   string
}

func (v Vendor) String() string {
	return fmt.Sprintf("Prefix: %x, Name: %s", v.Prefix, v.Name)
}

func (v Vendor) Contains(mac net.HardwareAddr) bool {
	return bytes.HasPrefix(mac, v.Prefix)
}

type Vendors []Vendor

func Lookup(mac net.HardwareAddr) Vendor {
	for _, vendor := range List {
		if vendor.Contains(mac) {
			return vendor
		}
	}
	return Vendor{Name: ""}
}
