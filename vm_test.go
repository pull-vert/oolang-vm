// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package oolang_vm

import (
	"encoding/binary"
	"oolang-vm/cpu"
	"testing"
)

// Test RunOnce a VM with a single HALT opcode instruction
func TestHaltOpcode(t *testing.T) {
	vcpu := cpu.NewVCPU(binary.LittleEndian)
	testBytes := []byte{0, 0, 0, 0}
	vcpu.Program = testBytes
	runOnce(vcpu)
	if vcpu.Pc != 1 {
		t.Errorf("expected pointer count to 1, was %d", vcpu.Pc)
	}
}

// Test Run a VM with a single LOAD opcode instruction
func TestLoadOpcode(t *testing.T) {
	vcpu := cpu.NewVCPU(binary.LittleEndian)
	testBytes := []byte{1, 0, 244, 1} // This is how we represent 500 using two u8s in little endian format
	vcpu.Program = testBytes
	run(vcpu)
	if vcpu.Pc != 4 {
		t.Errorf("expected pointer count to 4, was %d", vcpu.Pc)
	}
	if vcpu.Registers[0] != 500 {
		t.Errorf("expected Register 0 to contain 500, was %d", vcpu.Registers[0])
	}
}
