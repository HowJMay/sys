// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm64
// +build darwin
// +build !ios

package cpu

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

func osInit() {
	ARM64.HasFP = sysctlEnabled("hw.optional.floatingpoint")
	ARM64.HasASIMD = sysctlEnabled("hw.optional.neon")
	ARM64.HasCRC32 = sysctlEnabled("hw.optional.armv8_crc32")
	ARM64.HasATOMICS = sysctlEnabled("hw.optional.armv8_1_atomics")
	ARM64.HasFPHP = sysctlEnabled("hw.optional.neon_hpfp")
	ARM64.HasASIMDHP = sysctlEnabled("hw.optional.floatingpoint")
	ARM64.HasSHA3 = sysctlEnabled("hw.optional.armv8_2_sha3")
	ARM64.HasSHA512 = sysctlEnabled("hw.optional.armv8_2_sha512")
	ARM64.HasASIMDFHM = sysctlEnabled("hw.optional.armv8_2_fhm")

	// There are no hw.optional sysctl values for the below features on Mac OS 11.0
	// to detect their supported state dynamically. Assume the CPU features that
	// Apple Silicon M1 supports to be available as a minimal set of features
	// to all Go programs running on darwin/arm64.
	ARM64.HasEVTSTRM = true
	ARM64.HasAES = true
	ARM64.HasPMULL = true
	ARM64.HasSHA1 = true
	ARM64.HasSHA2 = true
	ARM64.HasCPUID = true
	ARM64.HasASIMDRDM = true
	ARM64.HasJSCVT = true
	ARM64.HasFCMA = true
	ARM64.HasLRCPC = true
	ARM64.HasDCPOP = true
	ARM64.HasSM3 = true
	ARM64.HasSM4 = true
	ARM64.HasASIMDDP = true
	ARM64.HasSVE = true
}

// The following is minimal copy of functionality from x/sys/unix so the cpu package can call
// sysctl without depending on x/sys/unix.

func sysctlEnabled(name string, args ...int) bool {
	mib, err := nametomib(name)
	if err != nil {
		return false
	}

	for _, a := range args {
		mib = append(mib, _C_int(a))
	}

	// Find size.
	n := uintptr(0)
	if err := sysctl(mib, nil, &n, nil, 0); err != nil {
		return false
	}

	return true
}

type _C_int int32

func sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) {
	var _zero uintptr
	var _p0 unsafe.Pointer
	if len(mib) > 0 {
		_p0 = unsafe.Pointer(&mib[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, errno := syscall.Syscall6(
		syscall.SYS___SYSCTL,
		uintptr(_p0),
		uintptr(len(mib)),
		uintptr(unsafe.Pointer(old)),
		uintptr(unsafe.Pointer(oldlen)),
		uintptr(unsafe.Pointer(new)),
		uintptr(newlen))
	if errno != 0 {
		return errno
	}
	return nil
}

// nametomib is a copy from "unix.nametomib()" in "unix/syscall_darwin.go".
func nametomib(name string) (mib []_C_int, err error) {
	const CTL_MAXNAME = 0xc
	const siz = unsafe.Sizeof(mib[0])

	// NOTE(rsc): It seems strange to set the buffer to have
	// size CTL_MAXNAME+2 but use only CTL_MAXNAME
	// as the size. I don't know why the +2 is here, but the
	// kernel uses +2 for its own implementation of this function.
	// I am scared that if we don't include the +2 here, the kernel
	// will silently write 2 words farther than we specify
	// and we'll get memory corruption.
	var buf [CTL_MAXNAME + 2]_C_int
	n := uintptr(CTL_MAXNAME) * siz

	p := (*byte)(unsafe.Pointer(&buf[0]))
	bytes, err := byteSliceFromString(name)
	if err != nil {
		return nil, err
	}

	// Magic sysctl: "setting" 0.3 to a string name
	// lets you read back the array of integers form.
	if err = sysctl([]_C_int{0, 3}, p, &n, &bytes[0], uintptr(len(name))); err != nil {
		return nil, err
	}
	return buf[0 : n/siz], nil
}

// byteSliceFromString is a simple copy of "unix.ByteSliceFromString()"
func byteSliceFromString(s string) ([]byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, fmt.Errorf("invalid argument in cpu.byteSliceFromString()")
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return a, nil
}
