/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"context"
	"testing"
	"time"

	"go.osspkg.com/casecheck"

	"go.arwos.org/arwos/sdk/rpc"
)

func TestIntegration_Controller_Hello(t *testing.T) {
	cli, err := rpc.NewClient("/tmp/app.sock", rpc.ClientMaxConns(1), rpc.ClientTimeout(time.Millisecond*50))
	casecheck.NoError(t, err)

	for i := 0; i < 100; i++ {
		req := rpc.NewRequest()
		req.SetMethod("com.example.simple.hello")
		resp := cli.Call(context.Background(), req)
		casecheck.NoError(t, resp.GetError())
		bb, err := resp.GetField(rpc.FieldBody)
		casecheck.NoError(t, err)
		casecheck.Equal(t, "Hello", string(bb))
	}
}

func Benchmark_Controller_Hello(b *testing.B) {
	cli, err := rpc.NewClient("/tmp/app.sock", rpc.ClientMaxConns(1))
	casecheck.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := rpc.NewRequest()
			req.SetMethod("com.example.simple.hello")
			resp := cli.Call(context.Background(), req)
			casecheck.NoError(b, resp.GetError())
			bb, err := resp.GetField(rpc.FieldBody)
			casecheck.NoError(b, err)
			casecheck.Equal(b, []byte("Hello"), bb)
		}
	})
}
