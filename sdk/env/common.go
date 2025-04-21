/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package env

import (
	"regexp"
	"strings"
)

var rex = regexp.MustCompile(`(?mUi)[^0-9a-z\_]+`)

func CanonicalName(s string) string {
	s = strings.TrimSpace(s)
	s = rex.ReplaceAllString(s, "")
	return strings.ToUpper(s)
}
