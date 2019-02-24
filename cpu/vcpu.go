package cpu

// A Virtual CPU
type VCPU struct {
	// Registers
	registers [32]uint64
}

//
// CPU / VM functions
//

// NewVM returns a new VCPU object.
func NewVCPU() *VCPU {
	x := &VCPU{}
	return x
}
