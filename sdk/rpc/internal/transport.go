/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"go.osspkg.com/ioutils/data"
)

type Transport struct {
	data *data.Buffer
	head *Header
}

func NewTransport() *Transport {
	return &Transport{
		data: data.NewBuffer(128),
		head: &Header{
			Code:   0,
			Method: "",
			Ctx:    make(map[string]string, 5),
			Sys:    make(map[SysCode]string, 1),
		},
	}
}

func (v *Transport) Reset() {
	v.head.Code = 0
	v.head.Method = ""
	v.SoftReset()
}

func (v *Transport) SoftReset() {
	for key := range v.head.Sys {
		delete(v.head.Sys, key)
	}
	for key := range v.head.Ctx {
		delete(v.head.Ctx, key)
	}
	v.data.Reset()
}

func (v *Transport) Code() int {
	return v.head.Code
}

func (v *Transport) SetCode(c int) {
	v.head.Code = c
}

func (v *Transport) GetError() error {
	if val, ok := v.head.Sys[SysError]; ok {
		return errors.New(val)
	}
	return nil
}

func (v *Transport) SetError(err error) {
	if err == nil {
		delete(v.head.Sys, SysError)
		return
	}
	v.head.Sys[SysError] = err.Error()
}

func (v *Transport) Deadline() time.Time {
	if val, ok := v.head.Sys[SysDeadline]; ok {
		t, err := time.Parse(time.RFC3339, val)
		if err == nil {
			return t
		}
	}
	return time.Time{}
}

func (v *Transport) SetDeadline(t time.Time) {
	v.head.Sys[SysDeadline] = t.Format(time.RFC3339)
}

func (v *Transport) Method() string {
	return v.head.Method
}

func (v *Transport) SetMethod(method string) {
	v.head.Method = method
}

func (v *Transport) GetCtx(key string) (string, bool) {
	key = strings.ToLower(key)
	val, ok := v.head.Ctx[key]
	return val, ok
}

func (v *Transport) GetCtxKeys() []string {
	keys := make([]string, 0, len(v.head.Ctx))
	for key := range v.head.Ctx {
		keys = append(keys, key)
	}
	return keys
}

func (v *Transport) SetCtx(key, val string) {
	key = strings.ToLower(key)
	v.head.Ctx[key] = val
}

func (v *Transport) Read(p []byte) (n int, err error) {
	return v.data.Read(p)
}

func (v *Transport) Bytes() []byte {
	return v.data.Bytes()
}

func (v *Transport) String() string {
	return v.data.String()
}

func (v *Transport) Write(p []byte) (n int, err error) {
	return v.data.Write(p)
}

func (v *Transport) Seek(offset int64, whence int) (int64, error) {
	return v.data.Seek(offset, whence)
}

func (v *Transport) Decode(r io.Reader) error {
	if _, err := v.data.ReadFrom(r); err != nil {
		return err
	}

	if _, err := v.data.Seek(-binary.MaxVarintLen64, data.SeekEnd); err != nil {
		return err
	}

	hl, _ := binary.Varint(v.data.Next(binary.MaxVarintLen64))

	if _, err := v.data.Seek(-binary.MaxVarintLen64-hl, data.SeekEnd); err != nil {
		return err
	}

	b := v.data.Next(int(hl))
	var h Header
	if err := json.Unmarshal(b, &h); err != nil {
		return err
	}
	v.head = &h

	v.data.Truncate(binary.MaxVarintLen64 + int(hl))

	if _, err := v.data.Seek(0, data.SeekStart); err != nil {
		return err
	}

	return nil
}

func (v *Transport) Encode(w io.Writer) error {
	if _, err := v.data.Seek(0, data.SeekEnd); err != nil {
		return err
	}

	b, err := json.Marshal(v.head)
	if err != nil {
		return err
	}

	n, err := v.data.Write(b)
	if err != nil {
		return err
	}

	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(buf, int64(n))
	if _, err := v.data.Write(buf); err != nil {
		return err
	}

	if _, err := v.data.Seek(0, data.SeekStart); err != nil {
		return err
	}

	if _, err := v.data.WriteTo(w); err != nil {
		return err
	}

	return nil
}
