package scripttest

import (
	"github.com/Autlin/miaospeed/interfaces"
	"github.com/Autlin/miaospeed/service/macros/script"
)

type ScriptTest struct {
	interfaces.ScriptTestDS
}

func (m *ScriptTest) Type() interfaces.SlaveRequestMatrixType {
	return interfaces.MatrixScriptTest
}

func (m *ScriptTest) MacroJob() interfaces.SlaveRequestMacroType {
	return interfaces.MacroScript
}

func (m *ScriptTest) Extract(entry interfaces.SlaveRequestMatrixEntry, macro interfaces.SlaveRequestMacro) {
	if mac, ok := macro.(*script.Script); ok {
		m.Key = entry.Params
		m.ScriptResult = mac.Store[entry.Params]
	}
}
