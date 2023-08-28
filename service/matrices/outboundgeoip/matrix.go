package outboundgeoip

import (
	"github.com/Autlin/miaospeed/interfaces"
	"github.com/Autlin/miaospeed/service/macros/geo"
)

type OutboundGeoIP struct {
	interfaces.OutboundGeoIPDS
}

func (m *OutboundGeoIP) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixOutboundGeoIP
}

func (m *OutboundGeoIP) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroGeo
}

func (m *OutboundGeoIP) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*geo.Geo); ok {
		m.MultiStacks = mac.OutStacks
	}
}
