/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal_test

import (
	"bytes"
	"fmt"
	"testing"

	"go.osspkg.com/casecheck"

	"go.arwos.org/arwos/sdk/rpc/internal"
)

func TestUnit_New(t *testing.T) {
	tr := internal.NewTransport()

	tr.SetCode(123)
	tr.SetMethod("com.example.app")
	tr.SetCtx("user", "1234")
	casecheck.NoError(t, tr.SetField("file1", []byte(`Hello world`)))
	casecheck.Error(t, tr.SetField("file1", []byte(`Hello world`)))

	buf := bytes.NewBuffer(nil)
	casecheck.NoError(t, tr.Encode(buf))

	fmt.Println(buf.String())

	tr = internal.NewTransport()
	casecheck.NoError(t, tr.Decode(buf))

	casecheck.Equal(t, 123, tr.Code())
	casecheck.Equal(t, "com.example.app", tr.Method())

	{
		val, ok := tr.GetCtx("user")
		casecheck.True(t, ok)
		casecheck.Equal(t, "1234", val)

		val, ok = tr.GetCtx("asd")
		casecheck.False(t, ok)
		casecheck.Equal(t, "", val)
	}

	{
		val, err := tr.GetField("file1")
		casecheck.NoError(t, err)
		casecheck.Equal(t, "Hello world", string(val))

		val, err = tr.GetField("file2")
		casecheck.Error(t, err)
		casecheck.Nil(t, val)
	}

}
