// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

import (
	"encoding/binary"
	"testing"
)

// Test create a new VCPU
func TestCreateVCPU(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	t.Logf("VCPU register length=%d, Program counter=%d, Program length=%d", len(vcpu.Registers), vcpu.Pc, len(vcpu.Program))
	if vcpu.Registers[0] != 0 {
		t.Errorf("New register contains a value!")
	}
	if len(vcpu.Registers) != 32 {
		t.Errorf("Register length is not 32 but %d", len(vcpu.Registers))
	}
	if vcpu.Pc != 0 {
		t.Errorf("Initial Program counter value must be 0, but is %d", vcpu.Pc)
	}
	if len(vcpu.Program) != 0 {
		t.Errorf("Program length is not 0 but %d", len(vcpu.Program))
	}
}

// Test Execute a single HALT opcode instruction
func TestExecuteHaltOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{0, 0, 0, 0}
	vcpu.Program = testBytes
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 1 {
		t.Errorf("expected pointer count to 1, was %d", vcpu.Pc)
	}
}

// Test Execute a single LOAD opcode instruction
func TestExecuteLoadOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{1, 0, 244, 1} // This is how we represent 500 using two u8s in little endian format
	vcpu.Program = testBytes
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[0] != 500 {
		t.Errorf("expected Register 0 to contain 500, was %d", vcpu.Registers[0])
	}
}

// Test Execute a single ADD opcode instruction
func TestExecuteAddOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{2, 0, 1, 2} // Register[2] = Register[0] + Register[1]
	vcpu.Program = testBytes
	vcpu.Registers[0] = 10
	vcpu.Registers[1] = 15
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[2] != 25 {
		t.Errorf("expected Register 2 to contain 25 (10 + 15), was %d", vcpu.Registers[2])
	}
}

// Test Execute a single SUB opcode instruction
func TestExecuteSubOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{3, 0, 1, 2} // Register[2] = Register[0] + Register[1]
	vcpu.Program = testBytes
	vcpu.Registers[0] = 10
	vcpu.Registers[1] = 15
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[2] != -5 {
		t.Errorf("expected Register 2 to contain -5 (10 - 15), was %d", vcpu.Registers[2])
	}
}

// Test Execute a single MUL opcode instruction
func TestExecuteMulOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{4, 0, 1, 2} // Register[2] = Register[0] + Register[1]
	vcpu.Program = testBytes
	vcpu.Registers[0] = 10
	vcpu.Registers[1] = 15
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[2] != 150 {
		t.Errorf("expected Register 2 to contain 150 (10 * 15), was %d", vcpu.Registers[2])
	}
}

// Test Execute a single DIV opcode instruction
func TestExecuteDivOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{5, 0, 1, 2} // Register[2] = Register[0] + Register[1]
	vcpu.Program = testBytes
	vcpu.Registers[0] = 10
	vcpu.Registers[1] = 15
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[2] != 0 {
		t.Errorf("expected Register 2 to contain 0 (10 / 15), was %d", vcpu.Registers[2])
	}
	if vcpu.remainder != 10 {
		t.Errorf("expected Register 2 to contain 10 (10 remainder 15), was %d", vcpu.remainder)
	}
}

// Test Execute a single DIV opcode instruction
func TestExecuteDiv2Opcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{5, 0, 1, 2} // Register[2] = Register[0] + Register[1]
	vcpu.Program = testBytes
	vcpu.Registers[0] = 21
	vcpu.Registers[1] = 10
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[2] != 2 {
		t.Errorf("expected Register 2 to contain 2 (21 / 10), was %d", vcpu.Registers[2])
	}
	if vcpu.remainder != 1 {
		t.Errorf("expected Register 2 to contain 1 (21 remainder 10), was %d", vcpu.remainder)
	}
}

// Test Execute a single Unknown opcode instruction
func TestExecuteUnknownOpcode(t *testing.T) {
	vcpu := NewVCPU(binary.LittleEndian)
	testBytes := []byte{200, 0, 0, 0}
	vcpu.Program = testBytes
	vcpu.ExecuteInstruction()
	if vcpu.Pc != 1 {
		t.Errorf("expected pointer count to 1, was %d", vcpu.Pc)
	}
}
