// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package oolang_vm

import (
	"oolang-vm/cpu"
)

//
// VM functions
//

// Loops as long as instructions can be executed.
func run(vcpu *cpu.VCPU) {
	isDone := false
	progLen := len(vcpu.Program)
	for !isDone {
		// If our program counter has exceeded the length of the program itself,
		// something has gone awry
		if vcpu.Pc >= progLen {
			break
		}
		isDone = vcpu.ExecuteInstruction()
	}
}

// Executes one instruction. Meant to allow for more controlled execution of the VM
func runOnce(vcpu *cpu.VCPU) {
	vcpu.ExecuteInstruction()
}
