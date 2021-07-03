// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const cacheLineSize = 64

func initOptions() {
	options = []option{
		{Name: "fp", Feature: &ARM64.HasFP},
		{Name: "asimd", Feature: &ARM64.HasASIMD},
		{Name: "evstrm", Feature: &ARM64.HasEVTSTRM},
		{Name: "aes", Feature: &ARM64.HasAES},
		{Name: "fphp", Feature: &ARM64.HasFPHP},
		{Name: "jscvt", Feature: &ARM64.HasJSCVT},
		{Name: "lrcpc", Feature: &ARM64.HasLRCPC},
		{Name: "pmull", Feature: &ARM64.HasPMULL},
		{Name: "sha1", Feature: &ARM64.HasSHA1},
		{Name: "sha2", Feature: &ARM64.HasSHA2},
		{Name: "sha3", Feature: &ARM64.HasSHA3},
		{Name: "sha512", Feature: &ARM64.HasSHA512},
		{Name: "sm3", Feature: &ARM64.HasSM3},
		{Name: "sm4", Feature: &ARM64.HasSM4},
		{Name: "sve", Feature: &ARM64.HasSVE},
		{Name: "crc32", Feature: &ARM64.HasCRC32},
		{Name: "atomics", Feature: &ARM64.HasATOMICS},
		{Name: "asimdhp", Feature: &ARM64.HasASIMDHP},
		{Name: "cpuid", Feature: &ARM64.HasCPUID},
		{Name: "asimrdm", Feature: &ARM64.HasASIMDRDM},
		{Name: "fcma", Feature: &ARM64.HasFCMA},
		{Name: "dcpop", Feature: &ARM64.HasDCPOP},
		{Name: "asimddp", Feature: &ARM64.HasASIMDDP},
		{Name: "asimdfhm", Feature: &ARM64.HasASIMDFHM},
	}
}

func archInit() {
	osInit()
}

// func getMIDR() uint64
