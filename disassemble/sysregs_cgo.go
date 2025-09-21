package disassemble

/*
#include "sysregs_gen.h"
#include "sysregs_fmt_gen.h"
*/
import "C"

// SystemReg wraps the C enum SystemReg
type SystemReg C.enum_SystemReg

// SystemReg constants - commonly used ones
const (
	SYSREG_NONE   = SystemReg(C.SYSREG_NONE)
	REG_EDSCR     = SystemReg(C.REG_EDSCR)
	REG_EDPRCR    = SystemReg(C.REG_EDPRCR)
	REG_UAOIMM    = SystemReg(C.REG_UAOIMM)
	REG_PANIMM    = SystemReg(C.REG_PANIMM)
	REG_SPSELIMM  = SystemReg(C.REG_SPSELIMM)
	REG_DITIMM    = SystemReg(C.REG_DITIMM)
	REG_SVCRIMM   = SystemReg(C.REG_SVCRIMM)
	REG_ICIALLUIS = SystemReg(C.REG_ICIALLUIS)
	REG_ICIALLU   = SystemReg(C.REG_ICIALLU)
)

// String returns the string representation of a system register
func (s SystemReg) String() string {
	// Use the C formatting function
	cStr := C.get_system_register_name(C.enum_SystemReg(s))
	if cStr != nil {
		return C.GoString(cStr)
	}
	return "UNKNOWN"
}

// IsValid checks if the system register is valid
func (s SystemReg) IsValid() bool {
	return s != SystemReg(C.SYSREG_NONE)
}
