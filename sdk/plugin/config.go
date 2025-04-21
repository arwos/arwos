/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func confWrite(w io.Writer, value string, args ...any) {
	if _, err := fmt.Fprintf(w, value+"\n\n", args...); err != nil {
		panic(err)
	}
}

func configBase(w io.Writer) {
	debug := strings.ToUpper(os.Getenv("DEBUG"))

	switch debug {
	case "DEV":
		confWrite(w, `
env: dev
log:
    file_path: /dev/stdout
    format: string
    level: 4
`)
	case "PROD":
		confWrite(w, `
env: prod
log:
    file_path: /dev/stdout
    format: syslog
    level: 3
`)
	}
}
