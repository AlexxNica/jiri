// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated via go generate.
// DO NOT UPDATE MANUALLY
package modules_and_v23

import "fmt"
import "testing"
import "os"

import "v.io/x/ref/lib/modules"
import "v.io/x/ref/lib/testutil"
import "v.io/x/ref/lib/testutil/v23tests"

func init() {
	modules.RegisterChild("modulesModulesAndV23Int", `Oh..`, modulesModulesAndV23Int)
}

func TestMain(m *testing.M) {
	testutil.Init()
	if modules.IsModulesProcess() {
		if err := modules.Dispatch(); err != nil {
			fmt.Fprintf(os.Stderr, "modules.Dispatch failed: %v\n", err)
			os.Exit(1)
		}
		return
	}
	cleanup := v23tests.UseSharedBinDir()
	r := m.Run()
	cleanup()
	os.Exit(r)
}
