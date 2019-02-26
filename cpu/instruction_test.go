// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

import "testing"

// Test HALT opcode
func TestHaltOpcode(t *testing.T) {
	if HALT != 0x00 {
		t.Errorf("HALT opcode byte value must be equal to 0x00!")
	}
}

// Test create a new Instruction
func TestCreateInstruction(t *testing.T) {
	instruction := NewInstruction(HALT)
	if instruction.opcode != HALT {
		t.Errorf("This insctruction must contain HALT opcode!")
	}
}
