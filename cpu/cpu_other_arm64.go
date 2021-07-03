// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !netbsd && !darwin && arm64
// +build !linux,!netbsd,!darwin,arm64

package cpu

// func doinit() {}

func osInit() {
	setMinimalFeatures()
}

// setMinimalFeatures fakes the minimal ARM64 features expected by
// TestARM64minimalFeatures.
func setMinimalFeatures() {
	ARM64.HasASIMD = true
	ARM64.HasFP = true
}
