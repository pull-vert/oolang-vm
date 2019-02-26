// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

// A Virtual CPU
type VCPU struct {
	// Registers
	Registers [32]uint64
	Pc        int
	Program   []byte
}

//
// CPU / VM functions
//

// Creates and returns a new Instruction
func NewVCPU() *VCPU {
	x := &VCPU{}
	return x
}

// Return the next Opcode
func (cpu *VCPU) NextOpcode() byte {
	opcode := cpu.Program[cpu.Pc]
	cpu.Pc += 1
	return opcode
}
