// make core cli
// input --> docker

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hey timbo")
	day := time.Now().YearDay()
	fmt.Println(day % 7)
}

/*
	cli features --

  a general purpose command line interface that will let you spawn a custom
	component architecture.

	step0 download 	-- provide multiple sources{curl, package managers, }
	step1 init 			-- install your self and your program name to the system
	step2 expose 		-- interface to the evm with this program a la desemmbler.
	step3 spawn 		-- allow this program/tool to communicate with itself on other comps
	step4 repead

  functions: cli() -
			flags:	os, db, privacy, network

*/

/*
	tools:
	disassembler
	debuger
	wrapper
	dameon
	rlp.encode rlp.decode





*/

//  this will get you the absolute day number with respect to the year.
//  day := time.Now().YearDay()
//
// if we mod this by seven and assume the first day of the week is sunday
// then we can infer the day of the week
// dayOfWeek := day % 7
//
// time.Now() returns 2017-08-22 12:19:53.368110759 -0500 CDT
