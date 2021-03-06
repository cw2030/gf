// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// go test *.go -bench=".*" -benchmem

package gutil

import (
    "testing"
)

func Benchmark_TryCatch(b *testing.B) {
    for i := 0; i < b.N; i++ {
        TryCatch(func() {
            
        }, func(err interface{}) {
            
        })
    }
}
