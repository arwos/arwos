/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

//go:generate easyjson

//easyjson:json
type Header struct {
	Code   int                `json:"i"`
	Method string             `json:"m"`
	Ctx    map[string]string  `json:"c,omitempty"`
	Sys    map[SysCode]string `json:"s,omitempty"`
}
