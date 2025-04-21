/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package rpc

import (
	"context"
	"io"
	"time"

	"go.osspkg.com/network/client"

	"go.arwos.org/arwos/sdk/rpc/errs"
	"go.arwos.org/arwos/sdk/rpc/internal"
)

type (
	opt struct {
		Timeout  time.Duration
		MaxConns uint64
	}
	Option func(o *opt)
)

func ClientTimeout(arg time.Duration) Option {
	return func(o *opt) {
		o.Timeout = arg
	}
}

func ClientMaxConns(arg uint64) Option {
	return func(o *opt) {
		o.MaxConns = arg
	}
}

type (
	Client interface {
		Call(ctx context.Context, req Request) Response
	}

	_client struct {
		address string
		cli     client.Client
		opt     opt
	}
)

func NewRequest() Request {
	return internal.NewTransport()
}

func NewClient(address string, opts ...Option) (Client, error) {
	var err error

	c := &_client{
		address: address,
		opt: opt{
			Timeout:  time.Second,
			MaxConns: 10,
		},
	}

	for _, option := range opts {
		option(&c.opt)
	}

	c.cli, err = client.New(client.Config{
		Network:  "unix",
		Address:  address,
		MaxConns: c.opt.MaxConns,
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *_client) Call(ctx context.Context, req Request) Response {
	resp := internal.NewTransport()

	tr, ok := req.(*internal.Transport)
	if !ok {
		resp.SetError(errs.ErrInvalidRequest)
		return resp
	}

	deadline := time.Now().Add(c.opt.Timeout)
	tr.SetDeadline(deadline)

	ctx, ctxCancel := context.WithDeadline(ctx, deadline)
	defer func() { ctxCancel() }()

	err := c.cli.Call(ctx, func(ctx context.Context, w io.Writer, r io.Reader) error {
		if err := tr.Encode(w); err != nil {
			return err
		}
		return resp.Decode(r)
	})
	if err != nil {
		resp.Reset()
		resp.SetMethod(tr.Method())
		resp.SetError(err)
	}

	return resp
}
