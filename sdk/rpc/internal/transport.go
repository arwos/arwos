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

	"go.arwos.org/arwos/sdk/rpc/errs"
)

type Transport struct {
	data *data.Buffer
	head *Header
}

func NewTransport() *Transport {
	return &Transport{
		data: data.NewBuffer(1024),
		head: &Header{
			Code:   0,
			Method: "",
			Fields: make([]Field, 0, 5),
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
	v.head.Fields = v.head.Fields[:0]
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
	return time.Now().Add(time.Second * 1)
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

func (v *Transport) SetField(name string, val []byte) (err error) {
	name = strings.ToLower(name)
	for _, field := range v.head.Fields {
		if field.Name == name {
			return errs.ErrAlreadyExist
		}
	}
	field := Field{Name: name, Pos: v.data.Size(), Len: 0}
	field.Len, err = v.data.Write(val)
	v.head.Fields = append(v.head.Fields, field)
	return
}

func (v *Transport) GetFields() []string {
	keys := make([]string, 0, len(v.head.Fields))
	for _, field := range v.head.Fields {
		keys = append(keys, field.Name)
	}
	return keys
}

func (v *Transport) GetField(name string) (val []byte, err error) {
	name = strings.ToLower(name)
	for _, field := range v.head.Fields {
		if strings.ToLower(field.Name) == name {
			if _, err = v.data.Seek(int64(field.Pos), data.SeekStart); err != nil {
				return
			}
			return v.data.Next(field.Len), nil
		}
	}

	return nil, errs.ErrNotFound
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
