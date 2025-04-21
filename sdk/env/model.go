/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package env

//go:generate easyjson

//easyjson:json
type Env struct {
	Description string `json:"d"`
	Key         string `json:"k"`
	Default     string `json:"v"`
}

//easyjson:json
type Envs []Env
