package vendors

import (
	"github.com/Autlin/miaospeed/interfaces"

	"github.com/Autlin/miaospeed/vendors/clash"
	"github.com/Autlin/miaospeed/vendors/invalid"
	"github.com/Autlin/miaospeed/vendors/local"
)

var registeredList = map[interfaces.VendorType]func() interfaces.Vendor{
	interfaces.VendorLocal: func() interfaces.Vendor {
		return &local.Local{}
	},
	interfaces.VendorClash: func() interfaces.Vendor {
		return &clash.Clash{}
	},
}

func Find(vendorType interfaces.VendorType) interfaces.Vendor {
	if vendor, ok := registeredList[vendorType]; ok {
		return vendor()
	}

	return &invalid.Invalid{}
}
