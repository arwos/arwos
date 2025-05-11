/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package manifest

//go:generate easyjson

const FileName = "manifest.yaml"

type Type string

const (
	TypeService   Type = "service"
	TypeComposite Type = "composite"
)

//easyjson:json
type Model struct {
	Type        Type    `json:"type" yaml:"type"`
	Path        string  `json:"path" yaml:"path"`
	Name        string  `json:"name" yaml:"name"`
	Package     string  `json:"package" yaml:"package"`
	Description string  `json:"description" yaml:"description"`
	Author      string  `json:"author" yaml:"author"`
	Version     string  `json:"version" yaml:"version"`
	Menu        Menu    `json:"menu" yaml:"menu"`
	Links       []Link  `json:"links,omitempty" yaml:"links,omitempty"`
	Alias       []Alias `json:"alias,omitempty" yaml:"alias,omitempty"`
}

//easyjson:json
type Link struct {
	Url         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
}

//easyjson:json
type Menu struct {
	Group string `json:"group" yaml:"group"`
	Title string `json:"title" yaml:"title"`
}

//easyjson:json
type Alias struct {
	Group  string `json:"group" yaml:"group"`
	Extend string `json:"extend" yaml:"extend"`
}
