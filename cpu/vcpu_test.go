package cpu

import "testing"

// Test a new register is empty
func TestCreateVCPU(t *testing.T) {
	vcpu := NewVCPU()
	if vcpu.registers[0] != 0 {
		t.Errorf("New register contains a value!")
	}
	t.Logf("Register's length is %d", len(vcpu.registers))
	if len(vcpu.registers) != 32 {
		t.Errorf("Register's length is not 32 but %d", len(vcpu.registers))
	}
}
