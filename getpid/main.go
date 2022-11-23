package main

import (
	"github.com/mitchellh/go-ps"
)

func main() {
	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, p := range list {
		if p.Executable() == "OneDrive.exe" {
			println(int(p.Pid()))
		}
		//log.Printf("Process %s with PID %d", p.Executable(), p.Pid())
	}
}

//	func getpid() {
//		processes, _ := process.Processes()
//		for _, process := range processes {
//			name, _ := process.Name()
//			pid, _ := process.Ppid()
//			if name == "explorer.exe" {
//				println(pid)
//			}
//
//		}
//	}
//func getpid() {
//	processes, _ := process.Processes()
//	for _, process := range processes {
//		name, _ := process.Name()
//		pid, _ := process.Ppid()
//		println(name, pid)
//
//	}
//}
//func main() {
//	getpid()
//}
