// Copyright 2019 oolang's author : Frédéric Montariol. Use of this source code is governed by the Apache 2.0 license.

package oolang_vm

import (
	"fmt"
	"oolang-vm/cpu"
)

//
// VM functions
//

func Run(vcpu *cpu.VCPU) {
	for {
		// If our program counter has exceeded the length of the program itself, something has
		// gone awry
		if vcpu.Pc >= len(vcpu.Program) {
			break
		}
		opcode := vcpu.NextOpcode()
		switch opcode {
		case byte(cpu.HALT):
			fmt.Printf("HALT encountered, terminating")
			return
		default:
			fmt.Printf("Unrecognized opcode found : %x ! Terminating!", opcode)
			return
		}
	}
}
