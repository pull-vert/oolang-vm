// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package oolang_vm

import (
	"oolang-vm/cpu"
	"testing"
)

// Test Run a VM with a single HALT opcode instruction
func TestOpcodeHalt(t *testing.T) {
	vcpu := cpu.NewVCPU()
	testBytes := []byte{0, 0, 0, 0}
	vcpu.Program = testBytes
	Run(vcpu)
	if vcpu.Pc != 1 {
		t.Errorf("expected pointer count to 1, was %d", vcpu.Pc)
	}
}

// Test Run a VM with a single Unknown opcode instruction
func TestOpcodeUnknown(t *testing.T) {
	vcpu := cpu.NewVCPU()
	testBytes := []byte{200, 0, 0, 0}
	vcpu.Program = testBytes
	Run(vcpu)
	if vcpu.Pc != 1 {
		t.Errorf("expected pointer count to 1, was %d", vcpu.Pc)
	}
}
