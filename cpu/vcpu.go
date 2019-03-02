// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

import (
	"encoding/binary"
	"fmt"
)

// A Virtual CPU
type VCPU struct {
	// Registers
	Registers [32]int64
	Pc        int
	Program   []byte
	endianess binary.ByteOrder
}

//
// CPU / VM functions
//

// Creates and returns a new Instruction
func NewVCPU(endianess binary.ByteOrder) *VCPU {
	x := &VCPU{}
	x.endianess = endianess
	return x
}

func (vcpu *VCPU) ExecuteInstruction() bool {
	opcodeByte := vcpu.NextByte()
	switch opcodeByte {
	case byte(HALT): // Normal terminate program
		fmt.Printf("HALT encountered, terminating")
		fmt.Println()
		return true
	case byte(LOAD): // load a uint16 (from operand 2 and 3) into Register[operand 1]
		register := vcpu.NextUint8()
		number := vcpu.NextUint16()
		vcpu.Registers[register] = int64(number) // Our registers are int64s, so we need to cast it
	case byte(ADD): // Register[operand 3] = Register[operand 1] + Register[operand 2]
		val1 := vcpu.Registers[vcpu.NextUint8()]
		val2 := vcpu.Registers[vcpu.NextUint8()]
		vcpu.Registers[vcpu.NextUint8()] = val1 + val2
	default: // Unrecognized opcode, we don't know what to do with it, terminate program
		fmt.Printf("Unrecognized opcodeByte found : %x ! Terminating!", opcodeByte)
		fmt.Println()
		return true
	}
	return false
}

// Return Program's next 8 bits
func (vcpu *VCPU) NextByte() byte {
	next8bit := vcpu.Program[vcpu.Pc]
	vcpu.Pc += 1
	return next8bit
}

// Return Program's next 8 bits as uint8
func (vcpu *VCPU) NextUint8() uint8 {
	return uint8(vcpu.NextByte())
}

// Return Program's next 16 bits as uint16
// todo use unsafe for better performance ?
func (vcpu *VCPU) NextUint16() uint16 {
	nextUint16 := vcpu.endianess.Uint16(vcpu.Program[vcpu.Pc:])
	vcpu.Pc += 2
	return nextUint16
}
