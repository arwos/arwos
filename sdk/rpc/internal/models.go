/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

//go:generate easyjson

//easyjson:json
type Header struct {
	Code   int                `json:"a"`
	Method string             `json:"b"`
	Fields []Field            `json:"c,omitempty"`
	Ctx    map[string]string  `json:"d,omitempty"`
	Sys    map[SysCode]string `json:"e,omitempty"`
}

//easyjson:json
type Field struct {
	Name string `json:"x"`
	Pos  int    `json:"y"`
	Len  int    `json:"z"`
}
