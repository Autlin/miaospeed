package rttping

import (
	"github.com/Autlin/miaospeed/interfaces"
	"github.com/Autlin/miaospeed/service/macros/ping"
)

type RTTPing struct {
	interfaces.RTTPingDS
}

func (m *RTTPing) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixRTTPing
}

func (m *RTTPing) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroPing
}

func (m *RTTPing) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*ping.Ping); ok {
		m.Value = mac.RTT
	}
}
