/*
		Date:					Sun Aug 27th
		Author: 			Tim Siwula
		File:					cli.go
		Purpose:			A command line interperter tool.
		Run:					go run cli.go run echo hello from from your cli.go
		pegasys:			go run cli.go pegasys
									go run cli.go pegasys docker


*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	//installPegasys()
	dockerize()
	// testing()

}

func installPegasys() {
	// pwd := exec.Command("pwd")
	// pwd.Run()
	// fmt.Printf("pegasys pwd: %v\n", pwd)



// to do: make sym link from:
// /usr/local/bin/pegasys
// to:
// $GOPATH/cli.go




// if on a mac:
// copy file to: /usr/local/bin/cli.go
// then get default shell:
// then add alias: echo alias pegasys="go run cli.go pegasys" >> ~/.zsh
// then update terminal: source ~/.zshrc
// now run pegasys

}

func dockerize() {

	switch os.Args[1] {
	case "pegasys":
			runDocker()
	// defualt:
	// 		panic("you did not enter the 'run' command. Currently that is all that is supported right now.")
	}

}

func runDocker() {
	fmt.Printf("pegasys will now run: %v\n", os.Args[2:])
	command := exec.Command(os.Args[2], os.Args[3:]...)
	// command := exec.Command("docker login")
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	//command.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWUTS,
	// }
	command.Run()
	//assert(command.Run())
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func testing() {
	fmt.Println("testing function")
	day := time.Now()
	fmt.Println(day)
}

/*
	cli features --

	pegasys.cli() arguments/flags:
			os,
			db,
			privacy,
			network,
			client-light?,
			vm,
			key management,

  a general purpose command line interface that will let you spawn a custom
	component architecture.

	step0 download 	-- provide multiple sources{curl, package managers, }
	step1 init 			-- install your self and your program name to the system
	step2 expose 		-- interface to the evm with this program a la desemmbler.
	step3 spawn 		-- allow this program/tool to communicate with itself on other comps
	step4 repead

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
