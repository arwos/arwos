/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package rpc

import (
	"context"
)

type Reader interface {
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	GetFields() []string
	GetField(name string) ([]byte, error)
}

type Writer interface {
	SetCode(c int)
	SetCtx(key, val string)
	SetField(name string, val []byte) error
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
	SetField(name string, val []byte) error
}

type Response interface {
	Code() int
	GetError() error
	GetCtx(key string) (string, bool)
	GetCtxKeys() []string
	GetFields() []string
	GetField(name string) ([]byte, error)
}
