// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

// Represents an opcode, which tells our interpreter what to do with the following operands
// an opcode is represented as a single byte (8-bits)
type Opcode byte

// list all opcodes
const (
	HALT Opcode = 0x00
	LOAD Opcode = 0x01
	ADD  Opcode = 0x02
)

// an instruction is 32bits, 8-bits for the opcode and up to 3 8-bits operands
type Instruction struct {
	opcode Opcode
}

//
// Opcode / Instructions functions
//

// Creates and returns a new Instruction
func NewInstruction(opcode Opcode) *Instruction {
	x := &Instruction{opcode: opcode}
	return x
}
