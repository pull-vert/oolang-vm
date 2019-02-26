// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package cpu

import "testing"

// Test create a new VCPU
func TestCreateVCPU(t *testing.T) {
	vcpu := NewVCPU()
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
