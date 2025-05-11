/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package env

//go:generate easyjson

const FileName = "envs.yaml"

//easyjson:json
type Env struct {
	Description string `json:"description" yaml:"description"`
	Key         string `json:"key" yaml:"key"`
	Default     string `json:"default" yaml:"default"`
}

//easyjson:json
type Model []Env
