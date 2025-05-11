/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package rpc

import (
	"context"
	"io"
	"time"
)

type Reader interface {
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	io.Reader
	io.Seeker
	Bytes() []byte
	String() string
}

type Writer interface {
	SetCode(c int)
	SetCtx(key, val string)
	io.Writer
	io.Seeker
}

type GuardCtx interface {
	Method() string
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	SetCtx(key, val string)
}

type (
	Handler func(ctx context.Context, w Writer, r Reader) error
	Guard   func(ctx context.Context, gc GuardCtx) error
)

type Server interface {
	AddHandler(method string, call Handler)
	AddGuard(call Guard, methods ...string)
}

type Request interface {
	SetMethod(method string)
	SetCtx(key, val string)
	SetDeadline(t time.Time)
	Deadline() time.Time
	io.ReadSeeker
	Bytes() []byte
	String() string
}

type Response interface {
	Code() int
	Method() string
	GetError() error
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	io.WriteSeeker
	Bytes() []byte
	String() string
}

type Transport interface {
	SoftReset()
	Code() int
	SetCode(c int)
	GetError() error
	SetError(err error)
	Deadline() time.Time
	SetDeadline(t time.Time)
	Method() string
	SetMethod(method string)
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	SetCtx(key, val string)
	io.ReadWriteSeeker
	Bytes() []byte
	String() string
}
