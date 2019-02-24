// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

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
